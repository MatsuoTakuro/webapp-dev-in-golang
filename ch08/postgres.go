package ch08

import (
	"log"
	"webapp-dev-in-golang/ch03"
)

func postgres() {
	ch03.ConnectToPgx()
	createAuthorsTable()
	createBooksTable()
	initializeAuthorsTable()
	initializeBooksTable()
}

func createAuthorsTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS authors(
		author_id varchar(32) NOT NULL,
		name varchar(100) NOT NULL,
		created_at timestamp with time zone,
		CONSTRAINT pk_authors PRIMARY KEY (author_id)
	)`
	_, err := ch03.PgDb.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func createBooksTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS books(
		id varchar(32) NOT NULL,
		title varchar(100) NOT NULL,
		author_id varchar(32) NOT NULL,
		created_at timestamp with time zone,
		CONSTRAINT pk_books PRIMARY KEY (id),
		CONSTRAINT fk_authors
      FOREIGN KEY(author_id)
			REFERENCES authors(author_id)
	)`
	_, err := ch03.PgDb.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeAuthorsTable() {
	// cmd := `TRUNCATE authors;`
	// _, err := ch03.Db.Exec(cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cmd := `INSERT INTO authors(
	// 	author_id,
	// 	name,
	// 	created_at
	// )
	// VALUES('0001', 'Shinizu-san',  '2020-07-10 00:00:00.000000+00'),
	// 			('0002', 'Shibukawa-san', 	'2020-07-11 00:00:00.000000+00')`
	// _, err := ch03.Db.Exec(cmd)
	// if err != nil {
	// 	log.Println(err)
	// }
}

func initializeBooksTable() {
	cmd := `TRUNCATE books;`
	_, err := ch03.PgDb.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd = `INSERT INTO books(
		id,
		title,
		author_id,
		created_at
	)
	VALUES('0000001', 'webapp development in golang', '0001', '2020-07-10 00:00:00.000000+00'),
				('0000002', 'practical golang', '0002', '2020-07-11 00:00:00.000000+00'),
				('0000003', 'system programming golang', '0002', '2020-07-11 00:00:00.000000+00')`
	_, err = ch03.PgDb.Exec(cmd)
	if err != nil {
		log.Println(err)
	}
}
