package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type WarehouseController struct {
	DB db.WarehouseDB
}

func (c WarehouseController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := c.DB.GetAll()
	_ = json.NewEncoder(w).Encode(items)
}

func (c WarehouseController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := c.DB.Get(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func (c WarehouseController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	checkError(err, "Created", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

func (c WarehouseController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := c.DB.Update(item)
	checkError(err, "Updated", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

func (c WarehouseController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := c.DB.Delete(id)
	checkError(err, "Deleted", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

