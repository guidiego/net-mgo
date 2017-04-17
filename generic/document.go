package generic

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Document struct {
	collection mgo.Collection
	Id bson.ObjectId `json:"id", bson:"_id"`
}

type MongoDocument interface {
	SaveDocument() error
	FindDocument() error
	UpdateDocument()
	DeleteDocument()
}

func (d *Document) SaveDocument() error {
	d.collection.Save(d)
	return nil
}

func (d *Document) FindDocument() error {
	d.collection.Find(bson.M{"_id": d.Id}).One(d)
	return nil
}

func (d *Document) UpdateDocument() {
	// return d.collection.Save(d)
}

func (d *Document) DeleteDocument() {
	// return d.collection.Save(d)
}

func NewDocument(collection string, idString string) Document {
	objectId := bson.NewObjectId()
	d := Document{ MgoDatabase.C(collection), objectId }

	if idString != nil {
		d.Id = objectId.ObjectIdHex(idString)
		d.FindDocument()
	}

	return d
}