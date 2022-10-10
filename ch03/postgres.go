package ch03

import (
	"database/sql"
	"log"
)

var PgDb *sql.DB
var DbErr error

type Repository struct {
	db *sql.DB
}

var configValues string = "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable"

func ConnectToPgx() {
	PgDb, DbErr = sql.Open("pgx", configValues)
	if nil != DbErr {
		log.Fatal("error when opening db", DbErr)
	}
}
func createUsersTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS users(
		user_id varchar(32) NOT NULL,
		user_name varchar(100) NOT NULL,
		created_at timestamp with time zone,
		CONSTRAINT pk_users PRIMARY KEY (user_id)
	)`
	_, err := PgDb.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeUsersTable() {
	cmd := `TRUNCATE users;`
	_, err := PgDb.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd = `INSERT INTO users(
		user_id,
		user_name,
		created_at
	)
	VALUES('0001', 'Gopher', '2020-07-10 00:00:00.000000+00'),
				('0002', 'Ferris', '2020-07-11 00:00:00.000000+00')`
	_, err = PgDb.Exec(cmd)
	if err != nil {
		log.Println(err)
	}
}

func createProductsTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS products(
		product_no integer NOT NULL,
		name varchar(100) NOT NULL,
		price integer NOT NULL,
		CONSTRAINT pk_products PRIMARY KEY (product_no)
	)`
	_, err := PgDb.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeProductsTable() {
	cmd := `TRUNCATE products;`
	_, err := PgDb.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd = `INSERT INTO products(
		product_no,
		name,
		price
	)
	VALUES(1, 'X', 1000),
				(2, 'Y', 2000)`
	_, err = PgDb.Exec(cmd)
	if err != nil {
		log.Println(err)
	}
}
