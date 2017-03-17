package generic

import (
	"github.com/gorilla/"
)

type GenericResource struct {
	Model MgoModel
	Repo GenericRepository
}

func (gr *GenericResource) List(w *http.ResponseWriter, r *http.Request) {
	go gr.Repo.List(r.URL.Query(), gr.Model.List)
}

func (gr *GenericResource) Detail(w *http.ResponseWriter, r *http.Request) {
	go gr.Repo.Detail(gr.Model.Unique)
}

func (gr *GenericResource) Create(w *http.ResponseWriter, r *http.Request) {
	go gr.Repo.Create(gr.Model.Unique)
}

func (gr *GenericResource) Update(w *http.ResponseWriter, r *http.Request) {
	go gr.Repo.Update(gr.Model.Unique)
}

func (gr *GenericResource) Delete(w *http.ResponseWriter, r *http.Request) {
	go gr.Repo.Delete(gr.Model.Unique)
}

func NewResource(model MgoModel, parser func(map[string][]string) map[string]interface{}, successCallback func(data interface{}), failCallback func([]MgoError)) {
	gr := GenericRepository{
		parser, successCallback, failCallback,
	}

	return GenericResource {
		model,
		gr,
	}
}

func NewGenericResource(model MgoModel) {
	gr := GenericRepository{
		GenericParser,
		GenericSuccessCallback,
		GenericFailCallback
	}

	return GenericResource {
		model,
		gr,
	}
}