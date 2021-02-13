package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/storage"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
)

type BrandController struct {
	DB storage.BrandDB
}

func (c BrandController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c BrandController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c BrandController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Brand
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c BrandController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Brand
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c BrandController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := c.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "delete")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

