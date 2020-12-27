package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
)

type CategoryDB struct {
	Ctx string // contexto, lugar, se usa para el log del error
	Query models.QueryDB // son los string de consulta a la BD
}

func (db CategoryDB) GetAll() ([]models.Category, error) {
	res := make([]models.Category, 0)
	var item models.Category

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetAll", db.Ctx, "Reading rows")
		return res, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			checkError(err, "GetAll", db.Ctx, "Scan rows")
			return res, err
		} else{
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res, nil
}

func (db CategoryDB) Get(id string) (models.Category, error) {
	var item models.Category

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetAll", db.Ctx, "Reading rows")
		return item, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			checkError(err, "Get", db.Ctx, "Scan rows")
			return item, err
		}
	}
	defer rows.Close()
	return item, nil
}


func (db CategoryDB) Create(item models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Update(id string, item models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
