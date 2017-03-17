package person

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/guidiego/generic-api-mgo/generic"
)

type MgoModel struct {
	Collection mgo.Collection
	Unique MgoInterface
	List MgoListInterface
}

func (mm *MgoModel) GetDocument(idString string) {
	id := bson.NewObjectId().ObjectIdHex(idString)
	mm.Collection.Find(bson.M{"_id": id}).One(mm.Unique)
}

func (mm *MgoModel) SaveDocument() MgoEditResponse, []MgoError {
	mm.Collection.Save(p)
}

func (mm *MgoModel) UpdateDocument() MgoEditResponse, []MgoError {
	return nil, nil
}

func (mm *MgoModel) DeleteDocument() MgoEditResponse, []MgoError {
	return nil, nil
}

func (mm *MgoModel) FindDocuments() []MgoError {
	mm.Collection.Find(bson.M{"_id": id}).Iter().All(pl.List.Results)
}

