package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"log"
)

func GetMovementsWarehouse(idWarehouse string) []models.Movement {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(queryMovement["listWarehouseId"].Q, idWarehouse)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
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

func GetMovements() []models.Movement {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(queryMovement["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
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

func GetMovement(id string) []models.Movement {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(queryMovement["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
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


func CreateMovement(item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryMovement["insert"].Q)
	fmt.Println(tsql)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWareHouse", item.IdWarehouse),
		sql.Named("DateTime", item.Date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func UpdateMovement(item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryMovement["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWareHouse", item.IdWarehouse),
		sql.Named("DateTime", item.Date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func DeleteMovement(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryMovement["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
