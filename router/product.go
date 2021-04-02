package router

import (
	"github.com/CarosDrean/api-amachay/controllers"
	"github.com/CarosDrean/api-amachay/db"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/query"
	"github.com/gorilla/mux"
)

func productRoutes(s *mux.Router)  {
	ctrl := controllers.ProductController{
		DB: db.ProductDB{Ctx: "Product DB", Query: query.Product},
	}
	s.HandleFunc("/all/{id}", mid.CheckSecurity(ctrl.GetAllStock)).Methods("GET")
	s.HandleFunc("/all-product-warehouse/{id}", mid.CheckSecurity(ctrl.GetProductWarehouse)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
	s.HandleFunc("/all-product-noignore/{idWarehouse}", mid.CheckSecurity(ctrl.GetAllNoIgnore)).Methods("GET")
	s.HandleFunc("/all-product-new/{idWarehouse}", mid.CheckSecurity(ctrl.GetAllNew)).Methods("GET")

}
	

