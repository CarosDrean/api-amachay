package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.ClientPerson, 0)

	clients := db.GetClients()
	for _, e := range clients {
		person := db.GetPerson(strconv.Itoa(int(e.IdPerson)))[0]
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

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetClient(id)
	var userPerson models.ClientPerson
	if len(items) > 0 {
		person := db.GetPerson(strconv.Itoa(int(items[0].IdPerson)))
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

func CreateClient(w http.ResponseWriter, r *http.Request) {
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
	idPerson, err := db.CreatePerson(person)
	checkError(err, "Created", "Person")
	client := models.Client{
		IdPerson: int(idPerson),
		Type:     item.Type,
	}
	result, err := db.CreateClient(client)
	checkError(err, "Created", "Client")

	_ = json.NewEncoder(w).Encode(result)
}
func UpdateClient(w http.ResponseWriter, r *http.Request) {
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

	result, err := db.UpdatePerson(person)

	client := models.Client{
		ID:       item.ID,
		Type:     item.Type,
		IdPerson: int(item.IdPerson),
	}

	result, err = db.UpdateClient(client)

	checkError(err, "Updated", "Client")
	_ = json.NewEncoder(w).Encode(result)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	client := db.GetClient(id)[0]
	result, err := db.DeletePerson(strconv.Itoa(int(client.IdPerson)))
	result, err = db.DeleteClient(id)
	checkError(err, "Deleted", "Client")

	_ = json.NewEncoder(w).Encode(result)
}
