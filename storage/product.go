package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"strconv"
)

type ProductDB struct {
	Ctx   string
	Query models.QueryDB
}

func (db ProductDB) GetAllStock(idWarehouse string) ([]models.Product, error) {
	res := make([]models.Product, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAllStock", idWarehouse)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db ProductDB) GetAll() ([]models.Product, error) {
	res := make([]models.Product, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAllStock", "")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db ProductDB) Get(id string) (models.Product, error) {
	res := make([]models.Product, 0)

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAllStock", "")
	if err != nil {
		return models.Product{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db ProductDB) Create(item models.Product) (int64, error) {
	ctx := context.Background()
	tsql := db.Query["insert"].Q + "select isNull(SCOPE_IDENTITY(),-1);"
	stmt, err := DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", item.Name),
		sql.Named("Description", item.Description),
		sql.Named("Price", item.Price),
		sql.Named("Stock", item.Stock),
		sql.Named("IdCategory", item.IdCategory),
		sql.Named("Perishable", item.Perishable))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

func (db ProductDB) Update(item models.Product) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Name", item.Name),
		sql.Named("Description", item.Description),
		sql.Named("Price", item.Price),
		sql.Named("Stock", item.Stock),
		sql.Named("IdCategory", item.IdCategory),
		sql.Named("Perishable", item.Perishable))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
func (db ProductDB) Delete(id string) (int64, error) {
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

func (db ProductDB) scan(rows *sql.Rows, err error, res *[]models.Product, ctx string,
	situation string, extra string) error {
	var item models.Product
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		var perishable sql.NullBool
		err := rows.Scan(&item.ID, &item.IdCategory, &item.Name, &item.Description, &item.Price, &item.Stock, &perishable)
		item.Perishable = perishable.Bool
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			if extra != "" {
				item.Stock = GetStock(extra, item.ID)
			}
			category, _ := CategoryDB{
				Ctx:   "Category storage",
				Query: query.Category,
			}.GetByID(strconv.Itoa(item.IdCategory))
			item.Category = category.Name
			*res = append(*res, item)
		}
	}
	return nil
}
