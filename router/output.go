package router

import (
	output "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func outputRoutes(s *mux.Router)  {
	s.HandleFunc("/", mid.CheckSecurity(output.GetOutputs)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(output.GetOutput)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(output.CreateOutput)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(output.UpdateOutput)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(output.DeleteOutput)).Methods("DELETE")
}
