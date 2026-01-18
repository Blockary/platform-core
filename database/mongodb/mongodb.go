package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Connect() {
	mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)

	c, err := mongo.Connect(clientOptions)
	if err != nil {
		return
	}

	client = c
}

func GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized")
	}

	database := client.Database(dbName)
	if database == nil {
		return nil, fmt.Errorf("failed to get database: %s", dbName)
	}

	collection := database.Collection(collectionName)
	if collection == nil {
		return nil, fmt.Errorf("failed to get collection: %s from database: %s", collectionName, dbName)
	}

	return collection, nil
}
