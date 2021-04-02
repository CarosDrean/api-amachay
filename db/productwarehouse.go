package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type ProductWarehouseDB struct{
	Ctx   string
	Query models.QueryDB
}

func (db ProductWarehouseDB) Get(idProduct string, idWarehouse string) (models.ProductWarehouse, error) {
	res := make([]models.ProductWarehouse, 0)

	tsql := fmt.Sprintf(query.ProductWarehouse["get"].Q, idProduct, idWarehouse)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "product warehouse", "Get")
	if err != nil {
		return models.ProductWarehouse{}, err
	}
	if len(res) == 0 {
		return models.ProductWarehouse{}, nil
	}
	defer rows.Close()
	return res[0], nil
}

func (db ProductWarehouseDB) Create(item models.ProductWarehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductWarehouse["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWarehouse", item.IdWarehouse),
		sql.Named("Ignore", item.Ignore))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ProductWarehouseDB) Update(id string, item models.ProductWarehouse) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.ProductWarehouse["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWarehouse", item.IdWarehouse),
		sql.Named("Ignore", item.Ignore))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}




func (db ProductWarehouseDB) scan(rows *sql.Rows, err error, res *[]models.ProductWarehouse, ctx string, situation string) error {
	var item models.ProductWarehouse
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Ignore)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}