package helper

import (
	"context"
	"fmt"

	"github.com/diganth-hm/libray/Backend/database"
	model "github.com/diganth-hm/libray/Backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddBook(book model.Book) {
	books, err := database.BookCollection.InsertOne(context.Background(), book)
	Err(err)
	fmt.Println("sucssfully added :", books)
}

func DeleteoneBook(bkname string) {
	deletecount, err := database.BookCollection.DeleteOne(context.Background(), bson.M{"name": bkname})
	if err == mongo.ErrNoDocuments {
		fmt.Println("No Book found in that name")
	} else {
		Err(err)
	}
	fmt.Println("deletcount is ", deletecount)
}

func DeleteAllBook() {
	deletecount, err := database.BookCollection.DeleteMany(context.Background(), bson.D{{}})
	Err(err)
	fmt.Println("Delete count for all students ", deletecount)
}

func DisplayAllBooks() []model.Book {
	cur, err := database.BookCollection.Find(context.Background(), bson.D{{}})
	Err(err)
	defer cur.Close(context.Background())
	var books []model.Book
	for cur.Next(context.Background()) {
		var book model.Book
		err = cur.Decode(&book)
		Err(err)
		books = append(books, book)
	}
	return books
}

func DisplayoneBook(name string) model.Book {
	var book model.Book
	result := database.BookCollection.FindOne(context.Background(), bson.M{"book": name})
	err := result.Decode(&book)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No Book found with the name :", name)
	} else {
		Err(err)
	}
	return book

}
