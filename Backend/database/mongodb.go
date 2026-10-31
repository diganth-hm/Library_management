package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionURL = "mongodb://localhost:27017"
const dbName = "library"

type MongoDB struct {
	Client              *mongo.Client
	StudentCollection   *mongo.Collection
	BookCollection      *mongo.Collection
	BorrowingCollection *mongo.Collection
}

func Connect() (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(connectionURL)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}
	//pinging mongodb to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)

	studentcollection := database.Collection("Student")
	bookcollection := database.Collection("Book")
	borrowingcollection := database.Collection("Borrow")

	db := &MongoDB{
		Client:              client,
		StudentCollection:   studentcollection,
		BookCollection:      bookcollection,
		BorrowingCollection: borrowingcollection,
	}
	return db, nil
}
