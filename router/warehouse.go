package router

import (
	warehouse "github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/gorilla/mux"
)

func warehouseRoutes(s *mux.Router) {
	s.HandleFunc("/", mid.CheckSecurity(warehouse.GetWarehouses)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(warehouse.GetWarehouse)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(warehouse.CreateWarehouse)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(warehouse.UpdateWarehouse)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(warehouse.DeleteWarehouse)).Methods("DELETE")
}
