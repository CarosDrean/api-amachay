package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ClientsController struct {
	DB db.ClientDB
	PersonDB db.PersonDB
}

func (c ClientsController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.ClientPerson, 0)

	clients, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	for _, e := range clients {
		person, _ := c.PersonDB.Get(strconv.Itoa(int(e.IdPerson)))

		item := models.ClientPerson{
			ID:       e.ID,
			IdPerson: int64(person.ID),
			Cel:      person.Cel,
			Type:     e.Type,
			Name:     person.Name,
			LastName: person.LastName,
			Phone:    person.Phone,
			Address:  person.Address,
			Dni:      person.Dni,
			Mail:     person.Mail,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (c ClientsController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	item, err := c.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}

	person, _ := c.PersonDB.Get(strconv.Itoa(int(item.IdPerson)))
	userPerson := models.ClientPerson{
		ID:       item.ID,
		IdPerson: int64(person.ID),
		Cel:      person.Cel,
		Type:     item.Type,
		Name:     person.Name,
		LastName: person.LastName,
		Phone:    person.Phone,
		Address:  person.Address,
		Dni:      person.Dni,
		Mail:     person.Mail,
	}
	_ = json.NewEncoder(w).Encode(userPerson)
}

func (c ClientsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.ClientPerson
	_ = json.NewDecoder(r.Body).Decode(&item)
	person := models.Person{
		Name:     item.Name,
		LastName: item.LastName,
		Cel:      item.Cel,
		Phone:    item.Phone,
		Address:  item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}
	idPerson, err := c.PersonDB.Create(person)
	if err != nil {
		returnErr(w, err, "crear person")
		return
	}
	client := models.Client{
		IdPerson: idPerson,
		Type:     item.Type,
	}
	result, err := c.DB.Create(client)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
func (c ClientsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.ClientPerson
	_ = json.NewDecoder(r.Body).Decode(&item)

	person := models.Person{
		ID:       int(item.IdPerson),
		Name:     item.Name,
		LastName: item.LastName,
		Phone:    item.Phone,
		Cel:      item.Cel,
		Address:  item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}

	result, err := c.PersonDB.Update(strconv.Itoa(int(item.IdPerson)), person)
	if err != nil {
		returnErr(w, err, "actualizar person")
		return
	}

	client := models.Client{
		ID:       item.ID,
		Type:     item.Type,
		IdPerson: item.IdPerson,
	}

	result, err = c.DB.Update(id, client)

	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (c ClientsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	client, _ := c.DB.Get(id)
	result, err := c.PersonDB.Delete(strconv.Itoa(int(client.IdPerson)))
	if err != nil {
		returnErr(w, err, "eliminar person")
		return
	}
	result, err = c.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
