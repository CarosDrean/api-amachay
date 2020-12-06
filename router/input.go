package router

import (
	input "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func inputRoutes(s *mux.Router)  {
	s.HandleFunc("/", mid.CheckSecurity(input.GetInputs)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(input.GetInput)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(input.CreateInput)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(input.UpdateInput)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(input.DeleteInput)).Methods("DELETE")
}
