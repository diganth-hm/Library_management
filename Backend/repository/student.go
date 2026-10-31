package repository

import (
	"context"
	"fmt"

	model "github.com/diganth-hm/libray/Backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StudentRepository struct {
	Collection *mongo.Collection
}

func NewStudentRepository(collection *mongo.Collection) *StudentRepository {
	return &StudentRepository{
		Collection: collection,
	}

}

func (r *StudentRepository) Addstu(student model.Student) error {
	stu, err := r.Collection.InsertOne(context.Background(), student)
	if err != nil {
		return err
	}
	fmt.Println("sucssfully added :", stu)
	return nil
}
func (r *StudentRepository) DisplayallStu() ([]model.Student, error) {
	cur, err := r.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var students []model.Student
	for cur.Next(context.Background()) {
		var student model.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		if err := cur.Err(); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *StudentRepository) Deleteallstu() error {
	deletecount, err := r.Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return err
	}
	fmt.Println("Delete count for all students ", deletecount)
	return nil
}

func (r *StudentRepository) Delonestu(usn string) (bool, error) {
	deletecount, err := r.Collection.DeleteOne(context.Background(), bson.M{"usn": usn})
	if err == mongo.ErrNoDocuments {
		fmt.Println("No student Found with the usn :", usn)
		return false, err
	} else {
		if err != nil {
			return false, err
		}
	}
	fmt.Println("succesfully deleted", deletecount)
	if deletecount.DeletedCount == 0 {
		return false, nil
	}
	return true, nil

}

func (r *StudentRepository) DisplayOnestu(usn string) (model.Student, error) {
	var student model.Student
	result := r.Collection.FindOne(context.Background(), bson.M{"usn": usn})
	err := result.Decode(&student)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No student Found with the usn :", usn)
		return model.Student{}, err
	} else {
		if err != nil {
			return model.Student{}, err
		}
	}
	return student, nil

}
