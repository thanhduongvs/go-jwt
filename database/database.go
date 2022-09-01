package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseURL  string = "mongodb://localhost:27017"
	databaseName string = "a7"
	DB           *mongo.Database
)

func Connect() error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB.")
	DB = client.Database("thanhduong")
	return nil
}

func Collection(collectionName string) *mongo.Collection {
	var collection *mongo.Collection = DB.Collection(collectionName)
	return collection
}
