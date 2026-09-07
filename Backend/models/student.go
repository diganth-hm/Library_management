package model

type Student struct {
	Name     string `json:"name" bson:"name"`
	Usn      string `json:"usn" bson:"usn"`
	Branch   string `json:"branch" bson:"branch"`
	Password string `json:"password" bson:"password"`
}
