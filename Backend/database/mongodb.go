package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionURL = "mongodb://localhost:27017"
const dbName = "library"

var (
	StudentCollection   *mongo.Collection
	BookCollection      *mongo.Collection
	BorrowingCollection *mongo.Collection
)

func init() {
	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//pinging mongodb to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(dbName)

	StudentCollection = database.Collection("Student")
	BookCollection = database.Collection("Book")
	BorrowingCollection = database.Collection("Borrow")
}
