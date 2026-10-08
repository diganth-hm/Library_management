package model

import "time"

type Borrowing struct {
	BorrowedAt string    `json:"borrowedat" bson:"borrowedat"`
	DueDate    time.Time `json:"duedate" bson:"duedate"`
	ReturnedAt time.Time `json:"returnedat" bson:"returnedat"`
	Status     bool      `json:"status" bson:"status"`
	Usn        string    `json:"usn" bson:"usn"`
	BookID     string    `json:"bookid" bson:"bookid"`
}
