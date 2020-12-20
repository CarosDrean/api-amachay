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

func GetProductsStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	idI, _ := strconv.Atoi(id)
	items := db.GetProductsStock(idI)
	_ = json.NewEncoder(w).Encode(items)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetProducts()
	_ = json.NewEncoder(w).Encode(items)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetProduct(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateProduct(item)
	checkError(err, "Created", "Product")

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	fmt.Println(id)
	fmt.Println(item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateProduct(item)
	checkError(err, "Updated", "Product")

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteProduct(id)
	checkError(err, "Deleted", "Product")

	_ = json.NewEncoder(w).Encode(result)
}

