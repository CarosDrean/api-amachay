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

type ProviderController struct {
	DB       db.ProviderDB
}

func (c ProviderController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.ProviderBusiness, 0)

	providers, err := c.DB.GetAll()
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener todos, error: %v", err))
		return
	}
	for _, e := range providers {
		business, _ := db.BusinessDB{}.Get(strconv.Itoa(int(e.IdBusiness)))
		item := models.ProviderBusiness{
			ID:         e.ID,
			IdBusiness: business.ID,
			Type:       e.Type,
			Name:       business.Name,
			RUC:        business.RUC,
			Address:    business.Address,
			Cel:        business.Cel,
			Phone:      business.Phone,
			Mail:       business.Mail,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (c ProviderController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.Get(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener, error: %v", err))
		return
	}
	business, _ := db.BusinessDB{}.Get(strconv.Itoa(int(item.IdBusiness)))
	providerBusiness := models.ProviderBusiness{
		ID:         item.ID,
		IdBusiness: business.ID,
		Type:       item.Type,
		Name:       business.Name,
		RUC:        business.RUC,
		Address:    business.Address,
		Cel:        business.Cel,
		Phone:      business.Phone,
		Mail:       business.Mail,
	}
	_ = json.NewEncoder(w).Encode(providerBusiness)
}

func (c ProviderController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var providerBusiness models.ProviderBusiness
	_ = json.NewDecoder(r.Body).Decode(&providerBusiness)
	business := models.Business{
		ID:      providerBusiness.IdBusiness,
		Name:    providerBusiness.Name,
		RUC:     providerBusiness.RUC,
		Address: providerBusiness.Address,
		Cel:     providerBusiness.Cel,
		Phone:   providerBusiness.Phone,
		Mail:    providerBusiness.Mail,
	}
	idBusiness, err := db.BusinessDB{}.Create(business)
	if err != nil || idBusiness == -1 {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear Business, error: %v", err))
		return
	}
	provider := models.Provider{
		ID:         providerBusiness.ID,
		IdBusiness: idBusiness,
		Type:       providerBusiness.Type,
	}
	result, err := c.DB.Create(provider)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
func (c ProviderController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var providerBusiness models.ProviderBusiness
	_ = json.NewDecoder(r.Body).Decode(&providerBusiness)
	business := models.Business{
		ID:      providerBusiness.IdBusiness,
		Name:    providerBusiness.Name,
		RUC:     providerBusiness.RUC,
		Address: providerBusiness.Address,
		Cel:     providerBusiness.Cel,
		Phone:   providerBusiness.Phone,
		Mail:    providerBusiness.Mail,
	}

	result, err := db.BusinessDB{}.Update(strconv.Itoa(int(providerBusiness.IdBusiness)), business)

	provider := models.Provider{
		ID:         providerBusiness.ID,
		IdBusiness: int64(providerBusiness.IdBusiness),
		Type:       providerBusiness.Type,
	}

	result, err = c.DB.Update(id, provider)

	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al actualizar, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (c ProviderController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	item, _ := c.DB.Get(id)
	result, err := db.BusinessDB{}.Delete(strconv.Itoa(int(item.IdBusiness)))
	result, err = c.DB.Delete(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al eliminar, error: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
