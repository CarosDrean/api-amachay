package router

import (
	user "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func userRoutes(s *mux.Router) {
	s.HandleFunc("/", mid.CheckSecurity(user.GetSystemUsers)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(user.GetSystemUser)).Methods("GET")
	s.HandleFunc("/", user.CreateSystemUser).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(user.UpdateSystemUser)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(user.DeleteSystemUser)).Methods("DELETE")
}
