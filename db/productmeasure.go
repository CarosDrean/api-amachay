package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type ProductMeasureDB struct{}

func (db ProductMeasureDB) GetAll() ([]models.ProductMeasure, error) {
	res := make([]models.ProductMeasure, 0)

	tsql := fmt.Sprintf(query.ProductMeasure["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "product measure", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db ProductMeasureDB) Get(id string) (models.ProductMeasure, error) {
	res := make([]models.ProductMeasure, 0)
	tsql := fmt.Sprintf(query.ProductMeasure["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "product measure", "Get")
	if err != nil {
		return models.ProductMeasure{}, err
	}
	if len(res) == 0 {
		return models.ProductMeasure{}, nil
	}
	defer rows.Close()
	return res[0], nil
}

func (db ProductMeasureDB) GetProduct(id string) (models.ProductMeasure, error) {
	res := make([]models.ProductMeasure, 0)
	tsql := fmt.Sprintf(query.ProductMeasure["getProduct"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "product measure", "Get")
	if err != nil {
		return models.ProductMeasure{}, err
	}
	if len(res) == 0 {
		return models.ProductMeasure{}, nil
	}
	defer rows.Close()
	return res[0], nil
}

func (db ProductMeasureDB) Create(item models.ProductMeasure) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductMeasure["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdMeasure", item.IdMeasure),
		sql.Named("Unity", item.Unity),
		sql.Named("MinAlert", item.MinAlert))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProductMeasureDB) Update(id string, item models.ProductMeasure) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductMeasure["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdMeasure", item.IdMeasure),
		sql.Named("Unity", item.Unity),
		sql.Named("MinAlert", item.MinAlert))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProductMeasureDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductMeasure["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProductMeasureDB) DeleteProduct(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductMeasure["deleteProduct"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProductMeasureDB) scan(rows *sql.Rows, err error, res *[]models.ProductMeasure, ctx string, situation string) error {
	var item models.ProductMeasure
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdMeasure, &item.Unity, &item.MinAlert)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
