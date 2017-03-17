package generic

type GenericRepository struct {
	parser func(map[string][]string) map[string]interface{}
	successCallback func(data interface{})
	failCallback func([]MgoError)
}

func (gr *GenericRepository) List(query map[string][]string, doc MgoListInterface) {
	mongoQuery := gr.parser(r.URL.Query())

	if errs := doc.FindDocument(mongoQuery); len(errs) > 0 {
		gr.failCallback(errs)
	}
	
	gr.successCallback(doc)
}

func (gr *GenericRepository) Detail(doc MgoInterface) {
	if errs := doc.GetDocument(); len(errs) > 0 {
		gr.failCallback(errs)
	}

	gr.successCallback(doc)
}

func (gr *GenericRepository) Create(doc MgoInterface) {
	if errs := doc.ValidDocument(); len(errs) > 0 {
		gr.failCallback(errs)
	}

	response, errs := doc.SaveDocument()

	if len(errs) > 0 {
		gr.failCallback(errs)
	}

	gr.successCallback(response)
}

func (gr *GenericRepository) Update(data interface{}) {


}

func (gr *GenericRepository) Delete(data interface{}) {

}