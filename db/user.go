package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func GetSystemUsers() []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Username, &item.Password, &item.Role, &item.IdWarehouse)
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func GetSystemUser(id string) []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows 1: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Username, &item.Password, &item.Role, &item.IdWarehouse)
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func CreateSystemUser(item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["insert"].Q)
	item.Password = encrypt(item.Password)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("UserName", item.Username),
		sql.Named("Password", item.Password),
		sql.Named("Rol", item.Role),
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("IdWarehouse", item.IdWarehouse))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func UpdateSystemUser(item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["update"].Q)

	user := GetSystemUser(strconv.Itoa(item.ID))[0]
	if user.Password != item.Password {
		item.Password = encrypt(item.Password)
	}

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("UserName", item.Username),
		sql.Named("Password", item.Password),
		sql.Named("Rol", item.Role),
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("IdWarehouse", item.IdWarehouse))

	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.RowsAffected()
}
func DeleteSystemUser(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func GetSystemUserFromUserName(userName string) []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["getUserName"].Q, userName)

	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Username, &item.Password, &item.Role, &item.IdWarehouse)
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func ValidateSystemUserLogin(user string, password string) (constants.State, string) {
	items := GetSystemUserFromUserName(user)
	if len(items) > 0 {
		if comparePassword(items[0].Password, password) {
			return constants.Accept, strconv.Itoa(items[0].ID)
		}
		return constants.InvalidCredentials, ""
	}
	return constants.NotFound, ""
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func encrypt(password string) string {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		return ""
	}
	return string(hashedPassword)
}
