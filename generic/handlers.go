package generic

import (
	"net/http"
	"github.com/gorilla/mux"
)

type RouterList struct {
	Path string
	List []Router
}

type Router struct {
	Name string
	Methods string
	Path string
	HandlerFunc http.HandlerFunc
}

func (r *RouterList) InsertInto(router *mux.Router) {
	s := router.PathPrefix(h.Path).Subrouter()
	
	for _, handle := range r.List {
		s.HandleFunc(handle.Path, handle.HandlerFunc).Name(handle.Name).Methods(handle.Methods)
	}
}

func NewGenericRouter(collection string, unique MgoInterface, list MgoListInterface) RouterList {
	c := MgoDatabase().C(collection)
	mm := MgoModel{c, unique, list}
	resource := NewResource(mm)

	return RouterList{collection, []Router{
		Router{collection + "List", "GET", "/", resource.List},
		Router{collection + "Detail", "GET", "/{id}", resource.Detail},
		Router{collection + "Create", "POST", "/", resource.Create},
		Router{collection + "Update", "PUT", "/{id}", resource.Update},
		Router{collection + "Delete", "DELETE", "/{id}", resource.Delete},
	}}	
}