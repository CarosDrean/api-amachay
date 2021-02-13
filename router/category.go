package router

import (
	"github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/storage"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

func categoryRoutes(s *mux.Router) {
	ctrl := controllers.CategoryController{
		DB: storage.CategoryDB{},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}

func categoryRoutesEcho(g *echo.Group) {
	h := controllers.NewCategory(storage.CategoryDB{})
	g.GET("/", h.GetAll)
}
