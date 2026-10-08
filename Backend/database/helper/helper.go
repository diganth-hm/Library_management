package helper

import (
	"context"
	"fmt"

	"github.com/diganth-hm/libray/Backend/database"
	model "github.com/diganth-hm/libray/Backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// error handling
func Err(err error) {
	if err != nil {
		panic(err)
	}

}
func Addstu(student model.Student) {
	stu, err := database.StudentCollection.InsertOne(context.Background(), student)
	Err(err)
	fmt.Println("sucssfully added :", stu)
}
func DisplayallStu() []model.Student {
	cur, err := database.StudentCollection.Find(context.Background(), bson.D{{}})
	Err(err)
	defer cur.Close(context.Background())
	var students []model.Student
	for cur.Next(context.Background()) {
		var student model.Student
		err = cur.Decode(&student)
		Err(err)
		students = append(students, student)
	}
	return students
}

func Deleteallstu() {
	deletecount, err := database.StudentCollection.DeleteMany(context.Background(), bson.D{{}})
	Err(err)
	fmt.Println("Delete count for all students ", deletecount)
}

func Delonestu(usn string) {
	deletecount, err := database.StudentCollection.DeleteOne(context.Background(), bson.M{"usn": usn})
	if err == mongo.ErrNoDocuments {
		fmt.Println("No student Found with the usn :", usn)
	} else {
		Err(err)
	}
	fmt.Println("deletecount", deletecount)

}

func DisplayOnestu(usn string) model.Student {
	var student model.Student
	result := database.StudentCollection.FindOne(context.Background(), bson.M{"usn": usn})
	err := result.Decode(&student)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No student Found with the usn :", usn)
	} else {
		Err(err)
	}
	return student

}

//helper func for books

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

}
