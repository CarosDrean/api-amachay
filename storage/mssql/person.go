package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
)

type PersonDB struct {
	Ctx   string
	Query models.QueryDB
}

func (db PersonDB) GetAll() ([]models.Person, error) {
	res := make([]models.Person, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}

	defer rows.Close()
	return res, nil
}

func (db PersonDB) Get(id string) (models.Person, error) {
	res := make([]models.Person, 0)

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetByID")
	if err != nil {
		return models.Person{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db PersonDB) Create(item models.Person) (int64, error) {
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
		sql.Named("LastName", item.LastName),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Dni", item.Dni),
		sql.Named("Mail", item.Mail))

	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

func (db PersonDB) Update(id string, item models.Person) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name),
		sql.Named("LastName", item.LastName),
		sql.Named("Cel", item.Cel),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Dni", item.Dni),
		sql.Named("Mail", item.Mail))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db PersonDB) Delete(id string) (int64, error) {
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

func (db PersonDB) scan(rows *sql.Rows, err error, res *[]models.Person, ctx string, situation string) error {
	var item models.Person
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.LastName, &item.Cel, &item.Phone, &item.Address, &item.Dni, &item.Mail)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
