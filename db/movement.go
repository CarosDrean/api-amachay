package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"strconv"
	"time"
)

type MovementDB struct{
	Ctx string
	Query models.QueryDB
}

func (db MovementDB) GetAllWarehouseFilter(filter models.Filter) ([]models.Movement, error) {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(db.Query["listWarehouseFilter"].Q, filter.ID, filter.Type, filter.DateFrom, filter.DateTo)
	rows, err := DB.Query(tsql)
	if err != nil {
		checkError(err, "GetAllWarehouseFilter", db.Ctx, "Reading rows")
		return res, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
		if err != nil {
			checkError(err, "GetWarehouseFilter", db.Ctx, "Scan rows")
			return res, err
		} else{
			item.Product = ProductDB{}.Get(strconv.Itoa(item.IdProduct))[0].Name
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res, err
}

func (db MovementDB) GetAllWarehouse(idWarehouse string) ([]models.Movement, error) {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(db.Query["listWarehouseId"].Q, idWarehouse)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetAllWarehouse", db.Ctx, "Reading rows")
		return res, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
		if err != nil {
			checkError(err, "GetAllWarehouse", db.Ctx, "Scan rows")
			return res, err
		} else{
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res, err
}

func (db MovementDB) GetAll() ([]models.Movement, error) {
	res := make([]models.Movement, 0)
	var item models.Movement

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetAll", db.Ctx, "Reading rows")
		return res, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
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

func (db MovementDB) Get(id string) (models.Movement, error) {
	var item models.Movement

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "Get", db.Ctx, "Reading rows")
		return item, err
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &item.IdClient)
		if err != nil {
			checkError(err, "Get", db.Ctx, "Scan rows")
			return item, err
		}
	}
	defer rows.Close()
	return item, nil
}

func (db MovementDB) Create(item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	date, err := time.Parse(time.RFC3339, item.Date + "T05:00:00Z")
	checkError(err, "Create", db.Ctx, "Convert Date")

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWareHouse", item.IdWarehouse),
		sql.Named("DateTime", date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MovementDB) Update(id string, item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	date, err := time.Parse(time.RFC3339, item.Date + "T05:00:00Z")
	checkError(err, "Update", db.Ctx, "Convert Date")

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWareHouse", item.IdWarehouse),
		sql.Named("DateTime", date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MovementDB) Delete(id string) (int64, error) {
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

func GetStock(idWarehouse int, idProduct int) float64 {
	var item float64

	tsql := fmt.Sprintf(query.Movement["stock"].Q, idWarehouse, idProduct)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetStock", "Movement DB", "Reading rows")
		return 0
	}
	for rows.Next(){
		var stock sql.NullFloat64
		err := rows.Scan(&stock)
		item = stock.Float64
		if err != nil {
			checkError(err, "GetStock", "Movement DB", "Scan rows")
			return 0
		}
	}
	defer rows.Close()
	return item
}
