package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type MeasureDB struct {}

func (db MeasureDB) GetAll() ([]models.Measure, error) {
	res := make([]models.Measure, 0)

	tsql := fmt.Sprintf(query.Measure["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "measure", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db MeasureDB) Get(id string) (models.Measure, error) {
	res := make([]models.Measure, 0)
	tsql := fmt.Sprintf(query.Measure["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "measure", "GetByID")
	if err != nil {
		return models.Measure{}, err
	}
	if len(res) == 0 {
		return models.Measure{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db MeasureDB) Create(item models.Measure) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Measure["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MeasureDB) Update(id string, item models.Measure) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Measure["update"].Q)
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

func (db MeasureDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Measure["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MeasureDB) scan(rows *sql.Rows, err error, res *[]models.Measure, ctx string, situation string) error {
	var item models.Measure
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
