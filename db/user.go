package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type UserDB struct {
	Ctx string
}

func (db UserDB) GetAll() ([]models.SystemUser, error) {
	res := make([]models.SystemUser, 0)

	tsql := fmt.Sprintf(query.SystemUser["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db UserDB) Get(id string) (models.SystemUser, error) {
	res := make([]models.SystemUser, 0)

	tsql := fmt.Sprintf(query.SystemUser["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return models.SystemUser{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db UserDB) Create(item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.SystemUser["insert"].Q)

	item.Password = encrypt(item.Password)

	nameIdWarehouse := sql.Named("IdWarehouse", nil)
	if item.IdWarehouse != -1 {
		nameIdWarehouse = sql.Named("IdWarehouse", item.IdWarehouse)
	}
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("UserName", item.Username),
		sql.Named("Password", item.Password),
		sql.Named("Rol", item.Role),
		sql.Named("IdPerson", item.IdPerson),
		nameIdWarehouse)

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db UserDB) Update(id string, item models.SystemUser) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.SystemUser["update"].Q)

	user, _ := db.Get(strconv.Itoa(item.ID))
	if user.Password != item.Password {
		item.Password = encrypt(item.Password)
	}

	nameIdWarehouse := sql.Named("IdWarehouse", nil)
	if item.IdWarehouse != -1 {
		nameIdWarehouse = sql.Named("IdWarehouse", item.IdWarehouse)
	}
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("UserName", item.Username),
		sql.Named("Password", item.Password),
		sql.Named("Rol", item.Role),
		sql.Named("IdPerson", item.IdPerson),
		nameIdWarehouse)

	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.RowsAffected()
}
func (db UserDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.SystemUser["delete"].Q)
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

	tsql := fmt.Sprintf(query.SystemUser["getUserName"].Q, userName)

	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		var idWarehouse sql.NullString
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Username, &item.Password, &item.Role, &idWarehouse)
		item.IdWarehouse = -1
		if idWarehouse.Valid {
			item.IdWarehouse, _ = strconv.Atoi(idWarehouse.String)
		}
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

func (db UserDB) scan(rows *sql.Rows, err error, res *[]models.SystemUser, ctx string, situation string) error {
	var item models.SystemUser
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		var idWarehouse sql.NullInt64
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Username, &item.Password, &item.Role, &idWarehouse)
		item.IdWarehouse = int(idWarehouse.Int64)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
