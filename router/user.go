package router

import (
	user "github.com/CarosDrean/api-amachay/controllers"
	"github.com/gorilla/mux"
)

func userRoutes(s *mux.Router) {
	s.HandleFunc("/", user.GetSystemUsers).Methods("GET")
	s.HandleFunc("/{id}", user.GetSystemUser).Methods("GET")
	s.HandleFunc("/", user.CreateSystemUser).Methods("POST")
	s.HandleFunc("/{id}", user.UpdateSystemUser).Methods("PUT")
	s.HandleFunc("/{id}", user.DeleteSystemUser).Methods("DELETE")
}
