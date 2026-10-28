package model

import "time"

type Borrowing struct {
	BorrowedAt string    `json:"borrowdate" bson:"borrowdate"`
	DueDate    time.Time `json:"duedate" bson:"duedate"`
	ReturnedAt time.Time `json:"returndate" bson:"returndate"`
	Status     bool      `json:"status" bson:"status"`
	Usn        string    `json:"usn" bson:"usn"`
	BookID     string    `json:"bookid" bson:"bookid"`
}
