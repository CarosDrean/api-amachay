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
	res := make([]models.ProductFill, 0)
	var params = mux.Vars(r)
	id, _ := params["id"]
	products, err := c.DB.GetAllStock(id)
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	for _, e := range products {
		productMeasure, _ := db.ProductMeasureDB{}.GetProduct(strconv.Itoa(e.ID))
		measure, _ := db.MeasureDB{}.Get(strconv.Itoa(productMeasure.IdMeasure))
		item := models.ProductFill{
			ID:               e.ID,
			Name:             e.Name,
			Description:      e.Description,
			Price:            e.Price,
			Stock:            e.Stock,
			IdProductMeasure: productMeasure.ID,
			IdCategory:       e.IdCategory,
			IdMeasure:        productMeasure.IdMeasure,
			Unity:            productMeasure.Unity,
			MinAlert:         productMeasure.MinAlert,
			Measure:          measure.Name,
			Perishable:       e.Perishable,
			Category:         e.Category,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func(c ProductController) GetProductWarehouse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	idWarehouse, _ := params["idWarehouse"]

	items, err := c.DB.GetProductWarehouse(idWarehouse)
	if err != nil {
		returnErr(w, err, "GetProductWarehouse")
		return
	}
	_ = json.NewEncoder(w).Encode(items)



}

func (c ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.ProductFill, 0)
	products, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	for _, e := range products {
		productMeasure, _ := db.ProductMeasureDB{}.GetProduct(strconv.Itoa(e.ID))
		item := models.ProductFill{
			ID:               e.ID,
			Name:             e.Name,
			Description:      e.Description,
			Price:            e.Price,
			Stock:            e.Stock,
			IdProductMeasure: productMeasure.ID,
			IdCategory:       e.IdCategory,
			IdMeasure:        productMeasure.IdMeasure,
			Unity:            productMeasure.Unity,
			MinAlert:         productMeasure.MinAlert,
			Perishable:       e.Perishable,
			Category:         e.Category,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}
func (c ProductController) GetAllNoIgnore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	idWarehouse, _ := params["idWarehouse"]

	res, err := db.ProductDB{}.GetAllNoIgnore(idWarehouse)

	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}
func (c ProductController) GetAllNew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	idWarehouse, _ := params["idWarehouse"]

	res, err := db.ProductDB{}.GetAllNew(idWarehouse)

	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}


func (c ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	productMeasure, err := db.ProductMeasureDB{}.GetProduct(id)
	if err != nil {
		returnErr(w, err, "obtener product measure")
		return
	}

	productFill := models.ProductFill{
		ID:               item.ID,
		Name:             item.Name,
		Description:      item.Description,
		Price:            item.Price,
		Stock:            item.Stock,
		IdProductMeasure: productMeasure.ID,
		IdCategory:       item.IdCategory,
		IdMeasure:        productMeasure.IdMeasure,
		Unity:            productMeasure.Unity,
		MinAlert:         productMeasure.MinAlert,
		Perishable:       item.Perishable,
	}

	_ = json.NewEncoder(w).Encode(productFill)
}

func (c ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.ProductFill
	_ = json.NewDecoder(r.Body).Decode(&item)
	product := models.Product{
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Stock:       0,
		IdCategory:  item.IdCategory,
		Perishable:  item.Perishable,
	}
	idProduct, err := c.DB.Create(product)
	if err != nil || idProduct == -1 {
		returnErr(w, err, "crear product")
		return
	}
	productMeasure := models.ProductMeasure{
		IdProduct: int(idProduct),
		IdMeasure: item.IdMeasure,
		Unity:     item.Unity,
		MinAlert:  item.MinAlert,
	}
	result, err := db.ProductMeasureDB{}.Create(productMeasure)
	if err != nil || idProduct == -1 {
		returnErr(w, err, "crear product measure")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.ProductFill
	_ = json.NewDecoder(r.Body).Decode(&item)
	product := models.Product{
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Stock:       0,
		IdCategory:  item.IdCategory,
		Perishable:  item.Perishable,
	}

	product.ID, _ = strconv.Atoi(id)
	result, err := c.DB.Update(product)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	productMeasure := models.ProductMeasure{
		ID:        item.IdProductMeasure,
		IdProduct: item.ID,
		IdMeasure: item.IdMeasure,
		Unity:     item.Unity,
		MinAlert:  item.MinAlert,
	}
	result, err = db.ProductMeasureDB{}.Update(strconv.Itoa(item.IdProductMeasure), productMeasure)
	if err != nil {
		returnErr(w, err, "update product measure")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.ProductMeasureDB{}.DeleteProduct(id)
	if err != nil {
		returnErr(w, err, "delete product measure")
		return
	}
	result, err = c.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

