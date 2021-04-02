package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
)

type WarehouseDB struct {
	Ctx   string
	Query models.QueryDB
}

func (db WarehouseDB) GetAll() ([]models.Warehouse, error) {
	res := make([]models.Warehouse, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db WarehouseDB) Get(id string) (models.Warehouse, error) {
	res := make([]models.Warehouse, 0)

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return models.Warehouse{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db WarehouseDB) Create(item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name),
		sql.Named("Address", item.Address),
		sql.Named("Ignore", item.State))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db WarehouseDB) Update(id string, item models.Warehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name),
		sql.Named("Address", item.Address),
		sql.Named("Ignore", item.State))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db WarehouseDB) Delete(id string) (int64, error) {
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

func (db WarehouseDB) scan(rows *sql.Rows, err error, res *[]models.Warehouse, ctx string, situation string) error {
	var item models.Warehouse
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Address, &item.State)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
