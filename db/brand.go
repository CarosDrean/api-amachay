package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type BrandDB struct {}

func (db BrandDB) GetAll() ([]models.Brand, error) {
	res := make([]models.Brand, 0)

	tsql := fmt.Sprintf(query.Brand["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "Brand", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db BrandDB) Get(id string) (models.Brand, error) {
	res := make([]models.Brand, 0)
	tsql := fmt.Sprintf(query.Brand["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "Brand", "Get")
	if err != nil {
		return models.Brand{}, err
	}
	if len(res) == 0 {
		return models.Brand{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db BrandDB) Create(item models.Brand) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Brand["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db BrandDB) Update(id string, item models.Brand) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Brand["update"].Q)
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

func (db BrandDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Brand["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db BrandDB) scan(rows *sql.Rows, err error, res *[]models.Brand, ctx string, situation string) error {
	var item models.Brand
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
