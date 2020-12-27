package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct {
	DB db.ProductDB
}

func (c ProductController) GetAllStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	items, err := c.DB.GetAllStock(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener todos con stock, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener todos, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c ProductController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)

	item.ID, _ = strconv.Atoi(id)
	result, err := c.DB.Update(item)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al actualizar, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c ProductController) Delete(w http.ResponseWriter, r *http.Request) {
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

