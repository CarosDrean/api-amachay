package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"log"
)

type ClientDB struct {}

func (db ClientDB) GetAll() []models.Client {
	res := make([]models.Client, 0)
	var item models.Client

	tsql := fmt.Sprintf(query.Client["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Type)
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

func (db ClientDB) Get(id string) []models.Client {
	res := make([]models.Client, 0)
	var item models.Client

	tsql := fmt.Sprintf(query.Client["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Type)
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


func (db ClientDB) Create(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ClientDB) Update(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ClientDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
