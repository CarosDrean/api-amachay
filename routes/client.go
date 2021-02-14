package routes

import (
	"github.com/CarosDrean/api-amachay/controllers"
	mid "github.com/CarosDrean/api-amachay/middleware"
	"github.com/CarosDrean/api-amachay/storage/mssql"
	"github.com/CarosDrean/api-amachay/storage/query-sql"
	"github.com/gorilla/mux"
)

func clientRoutes(s *mux.Router) {
	ctrl := controllers.ClientsController{
		DB: mssql.ClientDB{
			Ctx:   "Client mssql",
			Query: query_sql.Client,
		},
		PersonDB: mssql.PersonDB{Ctx: "PersonDB", Query: query_sql.Person},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}
