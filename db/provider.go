package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type ProviderDB struct {}

func (db ProviderDB) GetAll() ([]models.Provider, error) {
	res := make([]models.Provider, 0)

	tsql := fmt.Sprintf(query.Provider["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "provider", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db ProviderDB) Get(id string) (models.Provider, error) {
	res := make([]models.Provider, 0)
	tsql := fmt.Sprintf(query.Measure["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "provider", "Get")
	if err != nil {
		return models.Provider{}, err
	}
	if len(res) == 0 {
		return models.Provider{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db ProviderDB) Create(item models.Provider) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Provider["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdBusiness", item.IdBusiness),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProviderDB) Update(id string, item models.Provider) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Provider["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("IdBusiness", item.IdBusiness),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProviderDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Provider["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProviderDB) scan(rows *sql.Rows, err error, res *[]models.Provider, ctx string, situation string) error {
	var item models.Provider
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdBusiness, &item.Type)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}