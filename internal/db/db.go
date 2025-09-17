package db

import (
	"context"

	"api.link.henil.dev/internal/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

func Connect() *mongo.Client {
	uri := env.Get().MongoDBURI

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return c
}

func GetClient() *mongo.Client {
	if client == nil {
		client = Connect()
	}
	return client
}

func getDB() *mongo.Database {
	return GetClient().Database(env.Get().MongoDBName)
}

func init() {
	client = Connect()
}
