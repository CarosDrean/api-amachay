package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type CategoryDB struct {}

func (db CategoryDB) Create(item *models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Category["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Update(ID string, item *models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Category["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", ID),
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Delete(ID string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Category["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", ID))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) GetByID(ID string) (models.Category, error) {
	res := make([]models.Category, 0)
	tsql := fmt.Sprintf(query.Category["get"].Q, ID)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "db", "GetByID")
	if err != nil {
		return models.Category{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db CategoryDB) GetAll() ([]models.Category, error) {
	res := make([]models.Category, 0)

	tsql := fmt.Sprintf(query.Category["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "db", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db CategoryDB) scan(rows *sql.Rows, err error, res *[]models.Category, ctx string, situation string) error {
	var item models.Category
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
