package router

import (
	client "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func clientRoutes(s *mux.Router) {
	s.HandleFunc("/", mid.CheckSecurity(client.GetClients)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(client.GetClient)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(client.CreateClient)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(client.UpdateClient)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(client.DeleteClient)).Methods("DELETE")
}
