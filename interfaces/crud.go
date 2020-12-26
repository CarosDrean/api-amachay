package interfaces

import "net/http"

type CrudController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type CrudDB interface {
	GetAll() ([]interface{}, error)
	Get(id string) (interface{}, error)
	Create(i interface{}) (int64, error)
	Update(id int, i interface{}) (int64, error)
	Delete(i string) (int64, error)
}
