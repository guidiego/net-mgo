package generic

import (
	"os"	
	"gopkg.in/mgo.v2"
)

func MgoDatabase() mgo.Database {
	s, err := mgo.Dial(os.Getenv("MONGO_HOST"))

	if err != nil {
		panic(err)
	}

	return s.DB(os.Getenv("MONGO_DATABASE"))
}