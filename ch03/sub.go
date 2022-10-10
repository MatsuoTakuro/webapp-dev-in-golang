package ch03

import (
	"context"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func init() {
	ConnectToPgx()
	createUsersTable()
	initializeUsersTable()
	createProductsTable()
	initializeProductsTable()
}

func Sub() {
	defer PgDb.Close()
	list_3_3()
}

func list_3_3() {
	ctx := context.Background()
	r := Repository{
		db: PgDb,
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

	result1, err := tx.Exec("UPDATE users SET user_name=$1 WHERE user_id=$2", "Gopher senior", "0001")
	fmt.Println(result1.RowsAffected())
	if err != nil {
		return err
	}

	result2, err := tx.Exec("UPDATE users SET user_name=$1 WHERE user_id=$2", "Gopher senior bro", "0001")
	fmt.Println(result2.RowsAffected())
	if err != nil {
		return err
	}

	result3, err := tx.Exec("sql syntax err on purpose")
	fmt.Println(result3.RowsAffected())
	if err != nil {
		return err
	}

	return tx.Commit()
}
