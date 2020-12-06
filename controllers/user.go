package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetSystemUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.UserPerson, 0)

	users := db.GetSystemUsers()
	for _, e := range users {
		person := db.GetPerson(strconv.Itoa(int(e.IdPerson)))[0]
		item := models.UserPerson{
			ID:       e.ID,
			IdPerson: int64(person.ID),
			Cel:      person.Cel,
			Username: e.Username,
			Password: e.Password,
			Role:     e.Role,
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

func GetSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetSystemUser(id)
	var userPerson models.UserPerson
	if len(items) > 0 {
		person := db.GetPerson(strconv.Itoa(int(items[0].IdPerson)))
		userPerson = models.UserPerson{
			ID:       items[0].ID,
			IdPerson: int64(person[0].ID),
			Username: items[0].Username,
			Password: items[0].Password,
			Cel:      person[0].Cel,
			Role:     items[0].Role,
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

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userPerson models.UserPerson
	_ = json.NewDecoder(r.Body).Decode(&userPerson)
	person := models.Person{
		Name:     userPerson.Name,
		LastName: userPerson.LastName,
		Cel:      userPerson.Cel,
		Phone:    userPerson.Phone,
		Address:  userPerson.Address,
		Dni:      userPerson.Dni,
		Mail:     userPerson.Mail,
	}
	idPerson, err := db.CreatePerson(person)
	checkError(err, "Created", "Person")
	user := models.SystemUser{
		Username: userPerson.Username,
		Password: userPerson.Password,
		Role:     userPerson.Role,
		IdPerson: idPerson,
	}
	result, err := db.CreateSystemUser(user)
	checkError(err, "Created", "User")

	_ = json.NewEncoder(w).Encode(result)
}
func UpdateSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.UserPerson

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

	user := models.SystemUser{
		ID:       item.ID,
		Username: item.Username,
		Password: item.Password,
		Role:     item.Role,
		IdPerson: item.IdPerson,
	}

	result, err = db.UpdateSystemUser(user)

	checkError(err, "Updated", "User")
	_ = json.NewEncoder(w).Encode(result)
}

func DeleteSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	user := db.GetSystemUser(id)[0]
	result, err := db.DeletePerson(strconv.Itoa(int(user.IdPerson)))
	result, err = db.DeleteSystemUser(id)
	checkError(err, "Deleted", "User")

	_ = json.NewEncoder(w).Encode(result)
}
