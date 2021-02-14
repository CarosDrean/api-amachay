package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/storage/query-sql"
	"time"
)

type LotDB struct {}

func (db LotDB) GetAll() ([]models.Lot, error) {
	res := make([]models.Lot, 0)

	tsql := fmt.Sprintf(query_sql.Lot["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "lot", "GetAll")
	if err != nil  {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db LotDB) Get(id string) (models.Lot, error) {
	res := make([]models.Lot, 0)
	tsql := fmt.Sprintf(query_sql.Lot["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, "lot", "GetByID")
	if err != nil {
		return models.Lot{}, err
	}
	if len(res) == 0 {
		return models.Lot{}, nil
	}
	defer rows.Close()
	return res[0], nil
}


func (db LotDB) Create(item models.Lot) (int64, error) {
	ctx := context.Background()
	tsql := query_sql.Lot["insert"].Q + "select isNull(SCOPE_IDENTITY(),-1);"
	stmt, err := DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	dueDate := sql.Named("DueDate", nil)
	if item.DueDate != "" {
		date, err := time.Parse(time.RFC3339, item.DueDate+"T05:00:00Z")
		checkError(err, "Create", "lot mssql", "Convert Date")
		dueDate = sql.Named("DueDate", date)
	}

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Lot", item.Lot),
		dueDate)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

func (db LotDB) Update(id string, item models.Lot) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query_sql.Lot["update"].Q)
	dueDate := sql.Named("DueDate", nil)
	if item.DueDate != "" {
		date, err := time.Parse(time.RFC3339, item.DueDate+"T05:00:00Z")
		checkError(err, "Create", "lot mssql", "Convert Date")
		dueDate = sql.Named("DueDate", date)
	}
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Lot", item.Lot),
		dueDate)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db LotDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query_sql.Lot["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db LotDB) scan(rows *sql.Rows, err error, res *[]models.Lot, ctx string, situation string) error {
	var item models.Lot
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Lot, &item.DueDate)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
