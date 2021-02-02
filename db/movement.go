package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"github.com/CarosDrean/api-amachay/telegram"
	"strconv"
	"time"
)

type MovementDB struct {
	Ctx   string
	Query models.QueryDB
}

func (db MovementDB) GetAllWarehouseFilter(filter models.Filter) ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["listWarehouseFilter"].Q, filter.ID, filter.Type, filter.DateFrom, filter.DateTo)
	rows, err := DB.Query(tsql)
	err = db.scan(rows, err, &res, db.Ctx, "GetAllWarehouseFilter")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, err
}

func (db MovementDB) GetStockLot(idWarehouse int, idProduct int, idLot int) float64 {
	var item float64
	tsql := fmt.Sprintf(query.Movement["stockLot"].Q, idWarehouse, idProduct, idLot)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetStockLot", "Movement DB", "Reading rows")
		return 0
	}
	for rows.Next() {
		var stock sql.NullFloat64
		err := rows.Scan(&stock)
		item = stock.Float64
		if err != nil {
			checkError(err, "GetStockLot", "Movement DB", "Scan rows")
			return 0
		}
	}
	defer rows.Close()
	return item
}

func (db MovementDB) GetStockBrand(idWarehouse int, idProduct int, idBrand int) float64 {
	var item float64
	tsql := fmt.Sprintf(query.Movement["stockBrand"].Q, idWarehouse, idProduct, idBrand)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetStockBrand", "Movement DB", "Reading rows")
		return 0
	}
	for rows.Next() {
		var stock sql.NullFloat64
		err := rows.Scan(&stock)
		item = stock.Float64
		if err != nil {
			checkError(err, "GetStockLot", "Movement DB", "Scan rows")
			return 0
		}
	}
	defer rows.Close()
	return item
}

func (db MovementDB) GetAllLotsWarehouse(idProduct string, idWarehouse string)([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["getAllLotsWarehouse"].Q, idProduct, idWarehouse)
	rows, err := DB.Query(tsql)
	err = db.scan(rows, err, &res, db.Ctx, "GetAllLotsWarehouse")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, err
}

func (db MovementDB) GetAllWarehouse(idWarehouse string) ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["listWarehouseId"].Q, idWarehouse)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAllWarehouse")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, err
}

func (db MovementDB) GetAll() ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db MovementDB) Get(id string) (models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return models.Movement{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db MovementDB) Create(item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Create", db.Ctx, "Convert Date")

	idClient := sql.Named("IdClient", item.IdClient)
	idProvider := sql.Named("IdProvider", item.IdProvider)
	if item.Type == "input" {
		idClient = sql.Named("IdClient", nil)
	} else {
		idProvider = sql.Named("IdProvider", nil)
	}

	lot := sql.Named("IdLot", nil)
	product, _ := ProductDB{Ctx:   "Product DB", Query: query.Product}.Get(strconv.Itoa(item.IdProduct))
	if product.Perishable {
		lot = sql.Named("IdLot", item.IdLot)
	}

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWareHouse", item.IdWarehouse),
		sql.Named("DateTime", date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type),
		sql.Named("IdUser", item.IdUser),
		idClient,
		idProvider,
		lot,
		sql.Named("IdBrand", item.IdBrand))
	if err != nil {
		return -1, err
	}
	// obtener el stock y notificar a telegram
	GetDataForAlert(item)
	return result.RowsAffected()
}

func GetDataForAlert(item models.Movement) {
	stock := GetStock(strconv.Itoa(item.IdWarehouse), item.IdProduct)
	productMeasure, _ := ProductMeasureDB{}.GetProduct(strconv.Itoa(item.IdProduct))
	if stock <= float64(productMeasure.MinAlert) {
		product, _ := ProductDB{
			Ctx:   "For Alert",
			Query: query.Product,
		}.Get(strconv.Itoa(item.IdProduct))
		warehouse, _ := WarehouseDB{
			Ctx:   "For Alert",
			Query: query.Warehouse,
		}.Get(strconv.Itoa(item.IdWarehouse))
		measure, _ := MeasureDB{}.Get(strconv.Itoa(productMeasure.IdMeasure))
		telegram.AlertStock(warehouse.Name, product.Name, measure.Name, int(stock))
	}
}

func (db MovementDB) Update(id string, item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Update", db.Ctx, "Convert Date")

	idClient := sql.Named("IdClient", item.IdClient)
	idProvider := sql.Named("IdProvider", item.IdProvider)
	if item.Type == "input" {
		idClient = sql.Named("IdClient", nil)
	} else {
		idProvider = sql.Named("IdProvider", nil)
	}

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
		idClient,
		idProvider,
		sql.Named("IdLot", item.IdLot),
		sql.Named("IdBrand", item.IdBrand))
	if err != nil {
		return -1, err
	}
	GetDataForAlert(item)
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

func GetStock(idWarehouse string, idProduct int) float64 {
	var item float64

	tsql := fmt.Sprintf(query.Movement["stock"].Q, idWarehouse, idProduct)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetStock", "Movement DB", "Reading rows")
		return 0
	}
	for rows.Next() {
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

func (db MovementDB) scan(rows *sql.Rows, err error, res *[]models.Movement, ctx string, situation string) error {
	var item models.Movement
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		var idClient sql.NullInt64
		var idProvider sql.NullInt64
		var idLot sql.NullInt64
		var idBrand sql.NullInt64
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.Date, &item.Quantity, &item.Type,
			&item.IdUser, &idClient, &idProvider, &idLot, &idBrand)
		item.IdClient = int(idClient.Int64)
		item.IdProvider = int(idProvider.Int64)
		item.IdLot = int(idLot.Int64)
		item.IdBrand = int(idBrand.Int64)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			product, _ := ProductDB{Ctx:   "Movement", Query: query.Product}.Get(strconv.Itoa(item.IdProduct))
			productMeasure, _ := ProductMeasureDB{}.GetProduct(strconv.Itoa(product.ID))
			measure, _ := MeasureDB{}.Get(strconv.Itoa(productMeasure.IdMeasure))
			lot, _ := LotDB{}.Get(strconv.Itoa(item.IdLot))
			brand, _ := BrandDB{}.Get(strconv.Itoa(item.IdBrand))
			item.Lot = lot.Lot
			item.DueDate = lot.DueDate
			item.Brand = brand.Name
			item.Measure = measure.Name
			item.Product = product.Name
			item.Perishable = product.Perishable

			*res = append(*res, item)
		}
	}
	return nil
}
