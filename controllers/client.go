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

	clients := c.DB.GetAll()
	for _, e := range clients {
		person := c.PersonDB.Get(strconv.Itoa(int(e.IdPerson)))[0]
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

	items := c.DB.Get(id)
	var userPerson models.ClientPerson
	if len(items) > 0 {
		person := c.PersonDB.Get(strconv.Itoa(int(items[0].IdPerson)))
		userPerson = models.ClientPerson{
			ID:       items[0].ID,
			IdPerson: int64(person[0].ID),
			Cel:      person[0].Cel,
			Type:     items[0].Type,
			Name:     person[0].Name,
			LastName: person[0].LastName,
			Phone:    person[0].Phone,
			Address:  person[0].Address,
			Dni:      person[0].Dni,
			Mail:     person[0].Mail,
		}
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
	checkError(err, "Created", "Person")
	client := models.Client{
		IdPerson: idPerson,
		Type:     item.Type,
	}
	result, err := c.DB.Create(client)
	checkError(err, "Created", "Client")

	_ = json.NewEncoder(w).Encode(result)
}
func (c ClientsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.ClientPerson
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)

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

	result, err := c.PersonDB.Update(person)

	client := models.Client{
		ID:       item.ID,
		Type:     item.Type,
		IdPerson: item.IdPerson,
	}

	result, err = c.DB.Update(client)

	checkError(err, "Updated", "Client")
	_ = json.NewEncoder(w).Encode(result)
}

func (c ClientsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	client := c.DB.Get(id)[0]
	result, err := c.PersonDB.Delete(strconv.Itoa(int(client.IdPerson)))
	result, err = c.DB.Delete(id)
	checkError(err, "Deleted", "Client")

	_ = json.NewEncoder(w).Encode(result)
}
