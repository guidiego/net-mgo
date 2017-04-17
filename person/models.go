package person

import (
	"gopkg.in/mgo.v2/bson"

	. "github.com/guidiego/generic-api-mgo/generic"
)

func PersonCollection() mgo.Collection {
	return MgoDatabase().C("person")
}

type Person struct {
	Document
	Name string `json:"name", bson:"name"`
	Age int `json:"age", bson:"`
}

func New(name string, age int) {
	return Person{
		NewDocument("person", nil),
		name, age,
	}
}

func Get(idString string) {
	return Person{ NewDocument("person", idString) }
}

