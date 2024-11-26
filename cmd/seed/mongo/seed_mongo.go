package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"gopkg.in/yaml.v2"

	"github.com/tiborm/barefoot-bear/internal/seed/fileops"
	"github.com/tiborm/barefoot-bear/internal/seed/jsonops"
	"github.com/tiborm/barefoot-bear/internal/model"
)

// TODO check if you can use json model from cli code to validate the json file

var (
	MONGO_URI      string
	MONGO_DB       string
	FORCED_SEEDING bool
	config         Config
)

type Config struct {
	JsonFolder string `yaml:"json_folder"`

	CategoriesFolder     string `yaml:"categories_folder"`
	CategoriesCollection string `yaml:"categories_collection"`

	ProductsFolder     string `yaml:"products_folder"`
	ProductsCollection string `yaml:"products_collection"`

	InventoryFolder     string `yaml:"inventory_folder"`
	InventoryCollection string `yaml:"inventory_collection"`
}

func loadConfig(configFile string) Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return config
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	MONGO_URI = *flag.String("MONGO_URI", os.Getenv("MONGO_URI"), "MongoDB URI")
	MONGO_DB = *flag.String("MONGO_DB", os.Getenv("MONGO_DB"), "MongoDB Database Name")
	FORCED_SEEDING = *flag.Bool("FORCED_SEEDING", os.Getenv("FORCED_SEEDING") == "true", "Force seeding")
	flag.Parse()

	config = loadConfig("config.yaml")
}

func main() {
	categorySource := jsonops.JsonFolderToCollection[model.CategoryJsonResponse, model.Category]{
		JsonFolder: config.JsonFolder + config.CategoriesFolder,
		Collection: config.CategoriesCollection,
		Extractor: jsonops.ExtractCategories,
		IndexField: "id",
		JsonResponseStruct: &model.CategoryJsonResponse{},
		DataModel: &[]model.Category{},
	}

	productSource := jsonops.JsonFolderToCollection[model.ProductJsonResponse, model.Product]{
		JsonFolder: config.JsonFolder + config.ProductsFolder,
		Collection: config.ProductsCollection,
		Extractor: jsonops.ExtractProducts,
		IndexField: "id",
		JsonResponseStruct: &model.ProductJsonResponse{},
		DataModel: &[]model.Product{},
	}

	inventorySource := jsonops.JsonFolderToCollection[model.InventoryJsonResponse, model.Inventory]{
		JsonFolder: config.JsonFolder + config.InventoryFolder,
		Collection: config.InventoryCollection,
		Extractor: jsonops.ExtractInventory,
		IndexField: "",
		JsonResponseStruct: &model.InventoryJsonResponse{},
		DataModel: &[]model.Inventory{},
	}

	parseAndSeed(categorySource)
	parseAndSeed(productSource)
	parseAndSeed(inventorySource)

	fmt.Println("Seeding completed!")
}

func parseAndSeed[J jsonops.JsonResponseStruct, D jsonops.DataModel](source jsonops.JsonFolderToCollection[J, D]) {
	files := fileops.ReadFilenamesFromFolder(source.JsonFolder)

	var jsonExtract []*D

	for _, file := range files {
		jsonDoc, err := fileops.ReadJsonFile[J](source.JsonFolder + file)
		if err != nil {
			fmt.Printf("Failed to read JSON file: %s %v", file, err)
			continue
		}

		if source.Extractor != nil {
			jsonExtract = append(jsonExtract, source.Extractor(jsonDoc)...)
		}
	}

	dbClient := connectToDB()
	defer disconnectFromDB(dbClient)

	insertManyDocuments(dbClient, source.Collection, jsonExtract, source.IndexField)
}

func dereferenceSlice[D jsonops.DataModel](slice []*D) []D {
	var result []D
	for _, item := range slice {
		result = append(result, *item)
	}
	return result
}

func toInterfaceSlice[D any](typedSlice []*D) []interface{} {
	interfaceSlice := make([]interface{}, len(typedSlice))
	for i, v := range typedSlice {
		interfaceSlice[i] = *v
	}
	return interfaceSlice
}

func insertManyDocuments[D jsonops.DataModel](client *mongo.Client, collectionName string, documents []*D, indexField string) {
	collection := client.Database(MONGO_DB).Collection(collectionName)

	if FORCED_SEEDING {
		err := collection.Drop(context.Background())
		if err != nil {
			log.Fatalf("Failed to drop collection %s: %v", collectionName, err)
		}
		fmt.Printf("Collection %s dropped due to forced seeding\n", collectionName)
	}

	if indexField != "" {
		createIndexes(collection, indexField, true)
	}

	var dc int

	// If ordering false, the insertMany() method will not stop the insertion
	// of documents if duplicated items violate the unique index constraint.
	opts := options.InsertMany().SetOrdered(false)

	_, err := collection.InsertMany(context.Background(), toInterfaceSlice(documents), opts)
	if err != nil {
		var bulkWriteErrors mongo.BulkWriteException
		// Check if the error is a bulk write error like duplicate key
		// and treat it accordingly
		if errors.As(err, &bulkWriteErrors) {
			for _, writeError := range bulkWriteErrors.WriteErrors {
				fmt.Printf("Bulk write error: %v\n", writeError)
			}
			dc += len(bulkWriteErrors.WriteErrors)
		} else {
			// Any other error is fatal
			log.Fatalf("Failed to insert documents into MongoDB: %v\n", err)
		}
	}

	fmt.Printf("Documents count: %d\n", len(documents))
	fmt.Printf("Documents failed to insert: %d\n", dc)

	fmt.Printf("Documents successfully loaded into MongoDB collection: %s\n\n", collectionName)
}

func createIndexes(collection *mongo.Collection, fieldName string, isUnique bool) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: fieldName, Value: 1}},
		Options: options.Index().SetUnique(isUnique),
	}

	indexName, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		fmt.Printf("Failed to create index: %v\n", err)
	}

	fmt.Printf("Index created: %s\n", indexName)
}

func connectToDB() *mongo.Client {
	fmt.Printf("Connecting to MongoDB: %s\n", MONGO_URI)
	// Set up the MongoDB client options
	wc := writeconcern.W1()
	clientOptions := options.Client().ApplyURI(MONGO_URI).SetWriteConcern(wc)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB!\n\n")

	return client
}

func disconnectFromDB(client *mongo.Client) {
	// Disconnect from MongoDB
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB!")
}
