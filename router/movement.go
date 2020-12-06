package router

import (
	movement "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func movementRoutes(s *mux.Router) {
	s.HandleFunc("/all/{idWarehouse}", mid.CheckSecurity(movement.GetMovementsWarehouse)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(movement.GetMovements)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(movement.GetMovement)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(movement.CreateMovement)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(movement.UpdateMovement)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(movement.DeleteMovement)).Methods("DELETE")
}
