package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
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

type JsonFolderToCollection struct {
	JsonFolder string
	Collection string
	Extractor  func(*interface{}) []*interface{}
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
	dbClient := connectToDB()
	defer disconnectFromDB(dbClient)

	sources := []JsonFolderToCollection{
		{config.JsonFolder + config.CategoriesFolder, config.CategoriesCollection, extractCategories},
		{config.JsonFolder + config.ProductsFolder, config.ProductsCollection, extractProducts},
		{config.JsonFolder + config.InventoryFolder, config.InventoryCollection, extractInventory},
	}

	for _, source := range sources {
		files := readFilenamesFromFolder(source.JsonFolder)

		var jsonExtract []*interface{}

		for _, file := range files {
			jsonDoc, err := readJsonFile(source.JsonFolder + file)
			if err != nil {
				fmt.Printf("Failed to read JSON file: %s %v", file, err)
				continue
			}

			if source.Extractor != nil {
				jsonExtract = append(jsonExtract, source.Extractor(jsonDoc)...)
			}
		}

		insertManyDocuments(dbClient, source.Collection, jsonExtract)
	}

	fmt.Println("Seeding completed!")
}

func readFilenamesFromFolder(folderPath string) []string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames
}

func readJsonFile(jsonFile string) (*interface{}, error) {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s %v", jsonFile, err)
	}

	var result *interface{}
	if err := json.Unmarshal(file, &result); err != nil {
		return nil, fmt.Errorf("failed to decode: %s %v", jsonFile, err)
	}

	return result, nil
}

func extractCategories(categoriesJSON *interface{}) []*interface{} {
	var extract []*interface{}

	for _, c := range (*categoriesJSON).([]interface{}) {
		extract = append(extract, &c)
	}

	return extract
}

func extractProducts(productsByCategoryJSONs *interface{}) []*interface{} {
	var extract []*interface{}

	rs := (*productsByCategoryJSONs).(map[string]interface{})["results"].([]interface{})
	for _, r := range rs {
		is := r.(map[string]interface{})["items"].([]interface{})
		for _, pi := range is {
			product := pi.(map[string]interface{})["product"]
			if product != nil {
				extract = append(extract, &product)
			}
		}
	}

	return extract
}

func extractInventory(inventoryByProductJSON *interface{}) []*interface{} {
	var inventoryData []*interface{}

	for _, data := range (*inventoryByProductJSON).(map[string]interface{})["data"].([]interface{}) {
		inventoryData = append(inventoryData, &data)
	}

	return inventoryData
}

func dereferenceSlice(slice []*interface{}) []interface{} {
    var result []interface{}
    for _, item := range slice {
        result = append(result, *item)
    }
    return result
}

func insertManyDocuments(client *mongo.Client, collectionName string, documents []*interface{}) {
	collection := client.Database(MONGO_DB).Collection(collectionName)

	if FORCED_SEEDING {
		err := collection.Drop(context.Background())
		if err != nil {
			log.Fatalf("Failed to drop collection %s: %v", collectionName, err)
		}
		fmt.Printf("Collection %s dropped due to forced seeding\n", collectionName)
	}

	_, err := collection.InsertMany(context.Background(), dereferenceSlice(documents))
	if err != nil {
		log.Fatalf("Failed to insert documents into MongoDB: %v\n", err)
	}

	fmt.Printf("Documents successfully loaded into MongoDB collection: %s\n\n", collectionName)
}

func connectToDB() *mongo.Client {
	fmt.Printf("Connecting to MongoDB: %s\n", MONGO_URI)
	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI(MONGO_URI)

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
