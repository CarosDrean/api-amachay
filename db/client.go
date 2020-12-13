package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"log"
)

func GetClients() []models.Client {
	res := make([]models.Client, 0)
	var item models.Client

	tsql := fmt.Sprintf(queryClient["list"].Q)
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

func GetClient(id string) []models.Client {
	res := make([]models.Client, 0)
	var item models.Client

	tsql := fmt.Sprintf(queryClient["get"].Q, id)
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


func CreateClient(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryClient["insert"].Q)
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

func UpdateClient(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryClient["update"].Q)
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

func DeleteClient(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryClient["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
