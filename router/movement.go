package router

import (
	movement "github.com/CarosDrean/api-amachay/controllers"
	"github.com/CarosDrean/api-amachay/db"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/query"
	"github.com/gorilla/mux"
)

func movementRoutes(s *mux.Router) {
	ctrl := movement.MovementController{
		DB: db.MovementDB{
			Ctx:   "Movement DB",
			Query: query.Movement,
		},
	}
	s.HandleFunc("/all/{idWarehouse}", mid.CheckSecurity(ctrl.GetAllWarehouse)).Methods("GET")
	s.HandleFunc("/filter/", mid.CheckSecurity(ctrl.GetAllWarehouseFilter)).Methods("POST")
	s.HandleFunc("/all-lots-warehouse/", mid.CheckSecurity(ctrl.GetAllLotsWarehouse)).Methods("POST")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}
