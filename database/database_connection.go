package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	mongoDBURL := "mongodb://localhost:27017"
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, clientError := mongo.Connect(context, options.Client().ApplyURI(mongoDBURL))
	if clientError != nil {
		log.Fatal("Mongodb connection fail", clientError)
	}

	fmt.Println("Mongodb connected")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
