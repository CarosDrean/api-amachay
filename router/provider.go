package router

import (
	"github.com/CarosDrean/api-amachay/controllers"
	"github.com/CarosDrean/api-amachay/storage"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func providerRoutes(s *mux.Router) {
	ctrl := controllers.ProviderController{
		DB: storage.ProviderDB{},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}
