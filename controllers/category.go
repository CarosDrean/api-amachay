package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
)

type CategoryController struct {
	DB db.CategoryDB
}

func (c CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		// hay que probar imprimiendo err y err.Error()
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener todos, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.Get(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(item)
}

func (c CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, item)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al actualizar, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := c.DB.Delete(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al eliminar, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
