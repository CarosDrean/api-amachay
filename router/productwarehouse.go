package router

import (
	"github.com/CarosDrean/api-amachay/controllers"
	"github.com/CarosDrean/api-amachay/db"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/query"
	"github.com/gorilla/mux"
)

func productWarehouseRoutes(s *mux.Router) {
	ctrl := controllers.ProductWarehouseController{
		DB: db.ProductWarehouseDB{Ctx: "ProductWarehouse DB", Query: query.ProductWarehouse},
	}
	s.HandleFunc("/change-status/", mid.CheckSecurity(ctrl.ChangeStatus)).Methods("POST")

}
