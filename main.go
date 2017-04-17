package main

import (
	"github.com/guidiego/net-mgo/person"
)

func main() {
	person.New("Guilherme", 18);
	person.SaveDocument();

	println("Saved file!")
}