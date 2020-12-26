package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"log"
)

type WarehouseDB struct {}

func (db WarehouseDB) GetAll() []models.Warehouse {
	res := make([]models.Warehouse, 0)
	var item models.Warehouse

	tsql := fmt.Sprintf(query.Warehouse["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name, &item.Address, &item.State)
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

func (db WarehouseDB) Get(id string) []models.Warehouse {
	res := make([]models.Warehouse, 0)
	var item models.Warehouse

	tsql := fmt.Sprintf(query.Warehouse["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name, &item.Address, &item.State)
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


func (db WarehouseDB) Create(item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Warehouse["insert"].Q)
	fmt.Println(tsql)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name),
		sql.Named("Address", item.Address),
		sql.Named("State", item.State))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db WarehouseDB) Update(item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Warehouse["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Name", item.Name),
		sql.Named("Address", item.Address),
		sql.Named("State", item.State))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db WarehouseDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Warehouse["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
