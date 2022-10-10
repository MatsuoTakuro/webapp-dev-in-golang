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

func mysql() {
	MyDb = connectDB()
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
