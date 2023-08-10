package utils

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func InitDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("you must set your 'MONGODB_URI' environment variable")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// create new client and connect to server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	db = client.Database("xivpf")
	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}

func GetDBCollection(c string) *mongo.Collection {
	return db.Collection(c)
}
