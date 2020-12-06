package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetCategories()
	_ = json.NewEncoder(w).Encode(items)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetCategory(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateCategory(item)
	checkError(err, "Created", "Category")

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateCategory(item)
	checkError(err, "Updated", "Category")

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteCategory(id)
	checkError(err, "Deleted", "Category")

	_ = json.NewEncoder(w).Encode(result)
}
