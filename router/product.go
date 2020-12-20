package router

import (
	product "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func productRoutes(s *mux.Router)  {
	s.HandleFunc("/all/{id}", mid.CheckSecurity(product.GetProductsStock)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(product.GetProducts)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.GetProduct)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(product.CreateProduct)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.UpdateProduct)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.DeleteProduct)).Methods("DELETE")
}
	

