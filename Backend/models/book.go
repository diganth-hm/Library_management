package model

type Book struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}
