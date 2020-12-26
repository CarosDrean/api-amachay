package controllers

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/db"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	DB       db.UserDB
	PersonDB db.PersonDB
}

func (c UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.UserPerson, 0)

	users := c.DB.GetAll()
	for _, e := range users {
		person, _ := c.PersonDB.Get(strconv.Itoa(int(e.IdPerson)))
		item := models.UserPerson{
			ID:          e.ID,
			IdPerson:    int64(person.ID),
			Cel:         person.Cel,
			Username:    e.Username,
			Password:    e.Password,
			Role:        e.Role,
			Name:        person.Name,
			LastName:    person.LastName,
			Phone:       person.Phone,
			Address:     person.Address,
			Dni:         person.Dni,
			Mail:        person.Mail,
			IdWarehouse: e.IdWarehouse,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (c UserController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := c.DB.Get(id)
	var userPerson models.UserPerson
	if len(items) > 0 {
		person, _ := c.PersonDB.Get(strconv.Itoa(int(items[0].IdPerson)))
		userPerson = models.UserPerson{
			ID:          items[0].ID,
			IdPerson:    int64(person.ID),
			Username:    items[0].Username,
			Password:    items[0].Password,
			Cel:         person.Cel,
			Role:        items[0].Role,
			Name:        person.Name,
			LastName:    person.LastName,
			Phone:       person.Phone,
			Address:     person.Address,
			Dni:         person.Dni,
			Mail:        person.Mail,
			IdWarehouse: items[0].IdWarehouse,
		}
	}
	_ = json.NewEncoder(w).Encode(userPerson)
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
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
	idPerson, err := c.PersonDB.Create(person)
	checkError(err, "Created", "Person")
	user := models.SystemUser{
		Username:    userPerson.Username,
		Password:    userPerson.Password,
		Role:        userPerson.Role,
		IdPerson:    idPerson,
		IdWarehouse: userPerson.IdWarehouse,
	}
	result, err := c.DB.Create(user)
	checkError(err, "Created", "User")

	_ = json.NewEncoder(w).Encode(result)
}
func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
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

	result, err := c.PersonDB.Update(strconv.Itoa(int(item.IdPerson)), person)

	user := models.SystemUser{
		ID:          item.ID,
		Username:    item.Username,
		Password:    item.Password,
		Role:        item.Role,
		IdPerson:    item.IdPerson,
		IdWarehouse: item.IdWarehouse,
	}

	result, err = c.DB.Update(user)

	checkError(err, "Updated", "User")
	_ = json.NewEncoder(w).Encode(result)
}

func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	user := c.DB.Get(id)[0]
	result, err := c.PersonDB.Delete(strconv.Itoa(int(user.IdPerson)))
	result, err = c.DB.Delete(id)
	checkError(err, "Deleted", "User")

	_ = json.NewEncoder(w).Encode(result)
}
