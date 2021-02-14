package routes

import (
	"github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/storage/mssql"
	"github.com/gorilla/mux"
)

func invoiceRoutes(s *mux.Router) {
	ctrl := controllers.InvoiceController{
		DB: mssql.InvoiceDB{},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}

