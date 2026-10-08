package model

type Borrowing struct {
	BorrowedAt string `json:"borrowedat" bson:"borrowedat"`
	DueDate    string `json:"duedate" bson:"duedate"`
	ReturnedAt string `json:"returnedat" bson:"returnedat"`
	Status     bool   `json:"status" bson:"status"`
	Usn        string `json:"usn" bson:"usn"`
	BookID     int    ``
}
