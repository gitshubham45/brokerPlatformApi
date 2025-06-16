package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file %s \n", err)
	}
	mongoDbUri := os.Getenv("MONGO_URI")

	fmt.Printf("mongoUri : %s \n", mongoDbUri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		log.Fatal("error connecting to mongo db")
	}

	fmt.Println("connected to mongodb \\n")
	return client
}

var Client *mongo.Client = Init()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file %s \n", err)
	}
	
	dbName := os.Getenv("DB_NAME")
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}
