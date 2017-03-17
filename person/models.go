package person

import (
	"gopkg.in/mgo.v2/bson"

	. "github.com/guidiego/generic-api-mgo/generic"
)

func PersonCollection() mgo.Collection {
	return MgoDatabase().C("person")
}

type Person struct {
	Id bson.ObjectId `json:"id", bson:"_id"`
	Name string `json:"name", bson:"name"`
	Age int `json:"age", bson:"`
}

type PersonList struct {
	Results []Person
}

func (p *Person) GetDocument(idString string) {
	id := bson.NewObjectId().ObjectIdHex(idString)
	PersonCollection().Find(bson.M{"_id": id}).One(p)
}

func (p *Person) SaveDocument() MgoEditResponse, []MgoError {
	PersonCollection().Save(p)
}

func (p *Person) UpdateDocument() MgoEditResponse, []MgoError {
	return nil, nil
}

func (p *Person) DeleteDocument() MgoEditResponse, []MgoError {
	return nil, nil
}

func (pl *PersonList) FindDocuments() []MgoError {
	PersonCollection().Find(bson.M{"_id": id}).Iter().All(pl.Results)
}



