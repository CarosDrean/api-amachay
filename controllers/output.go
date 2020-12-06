package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetOutputs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetOutputs()
	_ = json.NewEncoder(w).Encode(items)
}

func GetOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetOutput(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func CreateOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Output
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateOutput(item)
	checkError(err, "Created", "Output")

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Output
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateOutput(item)
	checkError(err, "Updated", "Output")

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteOutput(id)
	checkError(err, "Deleted", "Output")

	_ = json.NewEncoder(w).Encode(result)
}


