package router

import (
	category "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func categoryRoutes(s *mux.Router) {
	s.HandleFunc("/", mid.CheckSecurity(category.GetCategories)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.GetCategory)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(category.CreateCategory)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.UpdateCategory)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.DeleteCategory)).Methods("DELETE")
}
