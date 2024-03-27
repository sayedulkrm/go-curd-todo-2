package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

	mongo_url := os.Getenv("MONGODB_URL")

	fmt.Println(mongo_url)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_url))

	if err != nil {
		logrus.Errorf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logrus.Errorf("Error pinging MongoDB: %v", err)
	}

	logrus.Info("Connected to MongoDB successfully. Connection")

	return client

}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Todo").Collection(collectionName)

	return collection

}
