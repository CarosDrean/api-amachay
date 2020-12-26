package controllers

import (
	"encoding/json"
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
	idI, _ := strconv.Atoi(id)
	items := c.DB.GetAllStock(idI)
	_ = json.NewEncoder(w).Encode(items)
}

func (c ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := c.DB.GetAll()
	_ = json.NewEncoder(w).Encode(items)
}

func (c ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := c.DB.Get(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func (c ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	checkError(err, "Created", "Product")

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
	checkError(err, "Updated", "Product")

	_ = json.NewEncoder(w).Encode(result)
}

func (c ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := c.DB.Delete(id)
	checkError(err, "Deleted", "Product")

	_ = json.NewEncoder(w).Encode(result)
}

