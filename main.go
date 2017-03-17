package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/guidiego/generic-api-mgo/person"
)

func main() {
	router := mux.NewRouter()
	for _, v :=  range []HandleList{
		person.HandleList,
	} { v.InsertInto(router) }
	

	http.ListenAndServe(":3000", router)
}