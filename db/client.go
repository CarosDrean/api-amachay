package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
)

type ClientDB struct {
	Ctx string
}

func (db ClientDB) GetAll() ([]models.Client, error) {
	res := make([]models.Client, 0)

	tsql := fmt.Sprintf(query.Client["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db ClientDB) Get(id string) (models.Client, error) {
	res := make([]models.Client, 0)

	tsql := fmt.Sprintf(query.Client["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return models.Client{}, err
	}
	defer rows.Close()
	return res[0], nil
}


func (db ClientDB) Create(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ClientDB) Update(item models.Client) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("IdPerson", item.IdPerson),
		sql.Named("Type", item.Type))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ClientDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(query.Client["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db ClientDB) scan(rows *sql.Rows, err error, res *[]models.Client, ctx string, situation string) error {
	var item models.Client
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdPerson, &item.Type)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
