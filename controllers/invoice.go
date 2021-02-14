package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/storage/mssql"
	"github.com/gorilla/mux"
	"net/http"
)

type InvoiceController struct {
	DB mssql.InvoiceDB
}

func (c InvoiceController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c InvoiceController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	_ = json.NewEncoder(w).Encode(item)
}

func (c InvoiceController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Invoice
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c InvoiceController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Invoice
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c InvoiceController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := c.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

