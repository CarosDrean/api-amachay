package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetSystemUsers() [] models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Username, &item.Rol, &item.Password,&item.IdPerson)
		if err != nil {
			log.Println(err)
			return res
		} else{
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
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson)
		if err != nil {
			log.Println(err)
			return res
		} else{
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func CreateSystemUser(item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Username", item.Username),
		sql.Named("password", item.Password),
		sql.Named("rol",item.Rol),
		sql.Named("IdPerson",item.IdPerson))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func UpdateSystemUser(item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["update"].Q)
	fmt.Println(tsql)
	fmt.Println(item)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Username", item.Username),
		sql.Named("password",item.Password),
		sql.Named("rol",item.Rol),
		sql.Named("IdPerson",item.IdPerson))

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
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson)
		if err != nil {
			log.Println(err)
			return res
		} else{
			res = append(res, item)
			log.Println(item.Password)
		}
	}
	defer rows.Close()
	return res
}

func ValidateSystemUserLogin(user string, password string) (constants.State, string){
	items := GetSystemUserFromUserName(user)
	if len(items) > 0 {
		if comparePassword(items[0].Password, password) {
			return constants.Accept, string(rune(items[0].ID))
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
