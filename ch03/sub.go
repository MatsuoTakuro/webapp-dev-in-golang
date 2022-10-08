package ch03

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB
var dbErr error

type Repository struct {
	db *sql.DB
}

var configValues string = "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable"

func connectToPgx() {
	db, dbErr = sql.Open("pgx", configValues)
	if nil != dbErr {
		log.Fatal(dbErr)
	}
}

func init() {
	connectToPgx()
}

func Sub() {
	list_3_3()
}

func list_3_3() {
	ctx := context.Background()
	r := Repository{
		db: db,
	}
	err := r.Update(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func (r *Repository) Update(ctx context.Context) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("update process 1")
	if err != nil {
		return err
	}

	_, err = tx.Exec("update process 2")
	if err != nil {
		return err
	}
	_, err = tx.Exec("update process 3")
	if err != nil {
		return err
	}

	return tx.Commit()
}
