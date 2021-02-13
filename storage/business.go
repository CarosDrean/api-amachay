package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type BusinessDB struct {}

func (db BusinessDB) GetAll() ([]models.Business, error) {
	res := make([]models.Business, 0)

	tsql := fmt.Sprintf(query.Business["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "measure", "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db BusinessDB) Get(id string) (models.Business, error) {
	res := make([]models.Business, 0)
	tsql := fmt.Sprintf(query.Business["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "business", "GetByID")
	if err != nil {
		return models.Business{}, err
	}
	if len(res) == 0 {
		return models.Business{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db BusinessDB) Create(item models.Business) (int64, error) {
	ctx := context.Background()
	tsql := query.Business["insert"].Q + "select isNull(SCOPE_IDENTITY(),-1);"
	stmt, err := DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", item.Name),
		sql.Named("RUC", item.RUC),
		sql.Named("Address", item.Address),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Mail", item.Mail))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

func (db BusinessDB) Update(id string, item models.Business) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Business["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name),
		sql.Named("RUC", item.RUC),
		sql.Named("Address", item.Address),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Mail", item.Mail))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db BusinessDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Business["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db BusinessDB) scan(rows *sql.Rows, err error, res *[]models.Business, ctx string, situation string) error {
	var item models.Business
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.RUC, &item.Address, &item.Cel, &item.Phone, &item.Mail)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}

