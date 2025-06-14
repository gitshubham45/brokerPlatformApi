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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error validating .env file")
	}

	mongoDbUri := os.Getenv("MONGO_URI")

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

func OpenCollection(client *mongo.Client , collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}