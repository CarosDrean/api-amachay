package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"net/http"
	"strconv"
)

type ProductWarehouseController struct {
	DB db.ProductWarehouseDB
}

func (pc ProductWarehouseController) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.ProductWarehouse
	_ = json.NewDecoder(r.Body).Decode(&item)
	productWarehouse, err := pc.DB.Get(strconv.Itoa(item.IdProduct), strconv.Itoa(item.IdWarehouse))
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	if productWarehouse.IdProduct == 0 {
		_, err = pc.DB.Create(item)
		if err != nil {
			returnErr(w, err, "crear")
			return
		}
		_ = json.NewEncoder(w).Encode("created")
		return
	}
	_, err = pc.DB.Update(strconv.Itoa(productWarehouse.ID), item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode("updated")
}

