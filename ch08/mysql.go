package ch08

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MyDb *sql.DB

type Repo struct {
	db *sql.DB
}

func NewRepo() *Repo {
	return &Repo{
		db: connectDB(),
	}
}

func connectToMySQL() {
	r := NewRepo()
	r.createAuthorsTable()
	r.createBooksTable()
	r.initializeAuthorsTable()
	r.initializeBooksTable()
}

func connectDB() *sql.DB {
	path := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true",
		"testuser", "pass", "testdb")

	return open(path, 100)
}

func open(path string, count uint) *sql.DB {
	MyDb, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = MyDb.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("mysql: db connected!!")
	return MyDb
}

func (r *Repo) createAuthorsTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS authors(
		author_id varchar(32) NOT NULL,
		name varchar(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
		PRIMARY KEY (author_id)
		)`
	_, err := r.db.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Repo) createBooksTable() {
	cmdU := `CREATE TABLE IF NOT EXISTS books(
			id varchar(32) NOT NULL,
			title varchar(100) NOT NULL,
			author_id varchar(32) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY(author_id)
			REFERENCES authors(author_id)
		)`
	_, err := r.db.Exec(cmdU)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Repo) initializeAuthorsTable() {
	// cmd := `TRUNCATE authors;`
	// _, err := r.db.Exec(cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cmd2 := `INSERT INTO authors(
		author_id,
		name
	)
	VALUES('0001', 'Shinizu-san'),
				('0002', 'Shibukawa-san')`
	_, err2 := r.db.Exec(cmd2)
	if err2 != nil {
		log.Println(err2)
	}
}

func (r *Repo) initializeBooksTable() {
	cmd := `TRUNCATE books;`
	_, err := r.db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cmd2 := `INSERT INTO books(
		id,
		title,
		author_id
	)
	VALUES('0000001', 'webapp development in golang', '0001'),
				('0000002', 'practical golang', '0002'),
				('0000003', 'system programming golang', '0002')`
	_, err2 := r.db.Exec(cmd2)
	if err2 != nil {
		log.Println(err2)
	}
}
