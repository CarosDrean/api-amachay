package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"log"
)

func GetInputs() []models.Input {
	res := make([]models.Input, 0)
	var item models.Input

	tsql := fmt.Sprintf(queryInput["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Date, &item.Quantity, &item.IdProduct)
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

func GetInput(id string) []models.Input {
	res := make([]models.Input, 0)
	var item models.Input

	tsql := fmt.Sprintf(queryInput["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Date, &item.Quantity, &item.IdProduct)
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

func CreateInput(item models.Input) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryInput["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Date", item.Date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("IdProduct", item.IdProduct))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
func UpdateInput(item models.Input) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryInput["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Date", item.Date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("IdProduct", item.IdProduct))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
func DeleteInput(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryInput["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
