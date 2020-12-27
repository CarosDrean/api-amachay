package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/CarosDrean/api-amachay/models"
)

type CategoryDB struct {
	Ctx string // contexto, lugar, se usa para el log del error
	Query models.QueryDB // son los string de consulta a la BD
}

func (db CategoryDB) GetAll() ([]models.Category, error) {
	res := make([]models.Category, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil
}

func (db CategoryDB) Get(id string) (models.Category, error) {
	res := make([]models.Category, 0)
	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "Get")
	if err != nil {
		return models.Category{}, err
	}
	defer rows.Close()
	return res[0], nil
}


func (db CategoryDB) Create(item models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Update(id string, item models.Category) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("Name", item.Name))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db CategoryDB) Delete(id string) (int64, error) {
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

func (db CategoryDB) scan(rows *sql.Rows, err error, res *[]models.Category, ctx string, situation string) error {
	var item models.Category
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			*res = append(*res, item)
		}
	}
	return nil
}
