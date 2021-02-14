package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/interfaces"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/storage/mssql"
	"github.com/CarosDrean/api-amachay/utils"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController struct {
	DB mssql.CategoryDB
}

type Category struct {
	storage interfaces.CategoryStorage
}

func NewCategory(storage interfaces.CategoryStorage) Category {
	return Category{storage}
}

func (cc Category) GetAll(c echo.Context) error {
	data, err := cc.storage.GetAll()
	if err != nil {
		response := utils.NewResponse(utils.Error, "Hubo un problema al obtener todas las categorias", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := utils.NewResponse(utils.Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

func (c CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (c CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.GetByID(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}

	_ = json.NewEncoder(w).Encode(item)
}

func (c CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Create(&item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Category
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := c.DB.Update(id, &item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func (c CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
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
