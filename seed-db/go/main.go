package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO implement force seeding

var MONGO_URI string

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	FORCED_SEEDING := os.Getenv("FORCED_SEEDING")
}

func main() {
	client := connectToDB()
	defer disconnectFromDB(client)

}

func loadCategoryJSON(client *mongo.Client) {

}

func connectToDB() *mongo.Client {
	log.Printf("Connectiong to MongoDB: %s", MONGO_URI)
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

	fmt.Println("Connected to MongoDB!")

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
