package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"log"
)

func GetWarehouses() []models.Warehouse {
	res := make([]models.Warehouse, 0)
	var item models.Warehouse

	tsql := fmt.Sprintf(queryWarehouse["list"].Q)
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

func GetWarehouse(id string) []models.Warehouse {
	res := make([]models.Warehouse, 0)
	var item models.Warehouse

	tsql := fmt.Sprintf(queryWarehouse["get"].Q, id)
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


func CreateWarehouse(item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryWarehouse["insert"].Q)
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

func UpdateWarehouse(item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryWarehouse["update"].Q)
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

func DeleteWarehouse(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryWarehouse["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
