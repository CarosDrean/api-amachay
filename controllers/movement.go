package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
)

type MovementController struct {
	DB db.MovementDB
}

func (c MovementController) GetInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Filter
	_ = json.NewDecoder(r.Body).Decode(&item)
	items, err := c.DB.GetInvoices(item.ID, item.AuxID)
	if err != nil {
		returnErr(w, err, "obtener todos warehouse filter")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c MovementController) GetAllWarehouseFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Filter
	_ = json.NewDecoder(r.Body).Decode(&item)

	items, err := c.DB.GetAllWarehouseFilter(item)
	if err != nil {
		returnErr(w, err, "obtener todos warehouse filter")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c MovementController) GetAllLotsWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Filter
	_ = json.NewDecoder(r.Body).Decode(&item)

	items, err := c.DB.GetAllLotsWarehouse(item.ID, item.AuxID)
	if err != nil {
		returnErr(w, err, "obtener todos lots warehouse")
		return
	}
	res := make([]models.Movement, 0)
	for i, e := range items {
		items[i].Quantity = float32(c.DB.GetStockLot(e.IdWarehouse, e.IdProduct, e.Lot))
		if items[i].Quantity > 0 {
			res = append(res, items[i])
		}
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (c MovementController) GetAllWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["idWarehouse"]

	items, err := c.DB.GetAllWarehouse(id)
	if err != nil {
		returnErr(w, err, "obtener todos warehouse")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c MovementController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c MovementController) Get(w http.ResponseWriter, r *http.Request) {
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

func (c MovementController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Movement
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c MovementController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Movement
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c MovementController) Delete(w http.ResponseWriter, r *http.Request) {
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


