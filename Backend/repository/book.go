package repository

import (
	"context"
	"fmt"

	model "github.com/diganth-hm/libray/Backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BookRepository struct {
	Collection *mongo.Collection
}

func NewBookRepository(collection *mongo.Collection) *BookRepository {
	return &BookRepository{
		Collection: collection,
	}
}

func (r *BookRepository) AddBook(book model.Book) error {
	books, err := r.Collection.InsertOne(context.Background(), book)
	if err != nil {
		return err
	}
	fmt.Println("sucssfully added :", books)
	return nil
}

func (r *BookRepository) DeleteoneBook(bkname string) (bool, error) {
	deletecount, err := r.Collection.DeleteOne(context.Background(), bson.M{"name": bkname})
	if err == mongo.ErrNoDocuments || err != nil {
		return false, err
	}
	fmt.Println("deletcount is ", deletecount)
	return true, nil
}

func (r *BookRepository) DeleteAllBook() error {
	deletecount, err := r.Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return err
	}
	fmt.Println("Delete count for all students ", deletecount)
	return nil
}

func (r *BookRepository) DisplayAllBooks() ([]model.Book, error) {
	cur, err := r.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var books []model.Book
	for cur.Next(context.Background()) {
		var book model.Book
		err = cur.Decode(&book)
		if err != nil {
			return nil, err
		}
		if err := cur.Err(); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepository) DisplayoneBook(name string) (model.Book, error) {
	var book model.Book
	result := r.Collection.FindOne(context.Background(), bson.M{"name": name})
	err := result.Decode(&book)
	if err == mongo.ErrNoDocuments || err != nil {
		return model.Book{}, err
	}
	return book, nil

}
