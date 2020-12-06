package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetWarehouses()
	_ = json.NewEncoder(w).Encode(items)
}

func GetWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetWarehouse(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateWarehouse(item)
	checkError(err, "Created", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateWarehouse(item)
	checkError(err, "Updated", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteWarehouse(id)
	checkError(err, "Deleted", "Warehouse")

	_ = json.NewEncoder(w).Encode(result)
}

