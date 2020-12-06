package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"log"
)

func GetOutputs() []models.Output {
	res := make([]models.Output, 0)
	var item models.Output

	tsql := fmt.Sprintf(queryOutput["list"].Q)
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

func GetOutput(id string) []models.Output {
	res := make([]models.Output, 0)
	var item models.Output

	tsql := fmt.Sprintf(queryOutput["get"].Q, id)
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

func CreateOutput(item models.Output) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryOutput["insert"].Q)
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

func UpdateOutput(item models.Output) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryOutput["update"].Q)
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
func DeleteOutput(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryOutput["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
