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

type InvoiceDB struct {}

func (db InvoiceDB) GetAll() ([]models.Invoice, error) {
	res := make([]models.Invoice, 0)

	tsql := fmt.Sprintf(query.Invoice["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "invoice", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db InvoiceDB) Get(id string) (models.Invoice, error) {
	res := make([]models.Invoice, 0)
	tsql := fmt.Sprintf(query.Invoice["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "invoice", "Get")
	if err != nil {
		return models.Invoice{}, err
	}
	if len(res) == 0 {
		return models.Invoice{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db InvoiceDB) Create(item models.Invoice) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Invoice["insert"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Update", "Create", "Convert Date")
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name),
		sql.Named("Code", item.Code),
		sql.Named("Date", date),
		sql.Named("IdImage", item.IdImage),
		sql.Named("IdProvider", item.IdProvider))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db InvoiceDB) Update(id string, item models.Invoice) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Invoice["update"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Update", "Update", "Convert Date")
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name),
		sql.Named("Code", item.Code),
		sql.Named("Date", date),
		sql.Named("IdImage", item.IdImage),
		sql.Named("IdProvider", item.IdProvider))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db InvoiceDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Invoice["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db InvoiceDB) scan(rows *sql.Rows, err error, res *[]models.Invoice, ctx string, situation string) error {
	var item models.Invoice
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		var idProvider sql.NullInt64
		err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.Date, &item.IdImage, &idProvider)
		item.IdProvider = int(idProvider.Int64)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			provider, _ := ProviderDB{}.Get(strconv.Itoa(item.IdProvider))
			business, _ := BusinessDB{}.Get(strconv.Itoa(int(provider.IdBusiness)))
			item.Provider = business.Name
			*res = append(*res, item)
		}
	}
	return nil
}
