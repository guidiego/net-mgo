package generic


import (
	"gopkg.in/mgo.v2"
	"os"
)

type MgoInterface interface{
	ValidDocument() []MgoError
	GetDocument() []MgoError
	SaveDocument() MgoEditResponse, []MgoError
	UpdateDocument() MgoEditResponse, []MgoError
	DeleteDocument() MgoEditResponse, []MgoError
}

type MgoListInterface interface{
	FindDocuments() []MgoError
}

type MgoError struct {
	Collection string `json:"collection"`
	Document bson.ObjectId `json:"document,omitempty"`
	Field string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type MgoEditResponse struct {
	Collection string `json:"collection"`
	DocumentIds []bson.ObjectId `json:"documents,omitempty"`
	Message string  `json:"message,omitempty"`
}

func MgoDatabase() mgo.Database {
	s, err := mgo.Dial(os.Getenv("MONGO_HOST"))

	if err != nil {
		panic(err)
	}

	return s.DB(os.Getenv("MONGO_DATABASE"))
}