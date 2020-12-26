package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"log"
)

type PersonDB struct {}

func (db PersonDB) GetAll() []models.Person {
	res := make([]models.Person, 0)
	var item models.Person

	tsql := fmt.Sprintf(query.Person["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.LastName, &item.Cel, &item.Phone, &item.Address, &item.Dni, &item.Mail)
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
func (db PersonDB) Get(id string) []models.Person {
	res := make([]models.Person, 0)
	var item models.Person

	tsql := fmt.Sprintf(query.Person["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.LastName, &item.Cel, &item.Phone, &item.Address, &item.Dni, &item.Mail)
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

func (db PersonDB) Create(item models.Person) (int64, error) {
	ctx := context.Background()
	tsql := query.Person["insert"].Q + "select isNull(SCOPE_IDENTITY(),-1);"

	stmt, err := DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", item.Name),
		sql.Named("LastName", item.LastName),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Dni", item.Dni),
		sql.Named("Mail", item.Mail))

	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

func (db PersonDB) Update(item models.Person) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Person["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Name", item.Name),
		sql.Named("LastName", item.LastName),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Dni", item.Dni),
		sql.Named("Mail", item.Mail))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db PersonDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Person["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
