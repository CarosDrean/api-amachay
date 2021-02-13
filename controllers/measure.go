package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/storage"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
)

type MeasureController struct {
	DB storage.MeasureDB
}

func (c MeasureController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c MeasureController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c MeasureController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Measure
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c MeasureController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Measure
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c MeasureController) Delete(w http.ResponseWriter, r *http.Request) {
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
