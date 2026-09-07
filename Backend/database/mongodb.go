package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionURL = "mongodb://localhost:27017"
const dbName = "library"
const colName = "student"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}

	Collection = client.Database(dbName).Collection(colName)
}
