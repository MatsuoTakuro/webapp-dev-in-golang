package ch08

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webapp-dev-in-golang/ch03"
)

func init() {
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
	_, err := ch03.Db.Exec(cmdU)
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
	_, err := ch03.Db.Exec(cmdU)
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
	_, err := ch03.Db.Exec(cmd)
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
	_, err = ch03.Db.Exec(cmd)
	if err != nil {
		log.Println(err)
	}
}

func Sub() {
	// list_8_3()
	list_8_6()
}

type AuthorID string

func (id AuthorID) Valid() bool {
	return false
}

type Author struct {
	AuthorID
	name string
}

func (a Author) Name() string {
	return a.name
}

type Book struct {
	ID    string
	Title string
	Author
}

func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		return nil, errors.New("GetAuthor: id is invalid")
	}

	return &Author{
			AuthorID: "0001",
			name:     "Taro",
		},
		nil
}

func GetAuthorName(b *Book) (string, error) {
	a, err := GetAuthor(b.AuthorID)
	if err != nil {
		return "", fmt.Errorf("GetAuthorName: %v", err)
	}
	return a.Name(), nil
}

func list_8_3() {
	b := &Book{
		ID:    "12345",
		Title: "XXXXX",
		Author: Author{
			AuthorID: "xxxxx",
			name:     "",
		},
	}
	_, err := GetAuthorName(b)
	if err != nil {
		log.Fatalln(err)
	}
}

func list_8_6() {
	// n, err := GetAuthorName2(context.Background(), "webapp development in golang")
	n, err := GetAuthorName2(context.Background(), "webapp development in golang2")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Author name:", n)
}

type Repository2 struct {
	db *sql.DB
}

func (r *Repository2) GetBook(ctx context.Context, title string) (*Book, error) {
	// - **ErrNoRows** is returned by Scan when **QueryRow** doesn't return a row.
	// - If **QueryRowContext** selects no rows, the *Row's Scan will return ErrNoRows.
	// - **ErrNoRows** can not be returned by Scan even if **Query** or **QueryContext** don't return a row. (err qual to nil get returned)
	// [sql](https://pkg.go.dev/database/sql#pkg-variables)
	// [Does db.Query return ErrNoRows?](https://stackoverflow.com/questions/60123848/does-db-query-return-errnorows)
	row := r.db.QueryRowContext(ctx, "SELECT id, title, author_id FROM books WHERE title=$1", title)

	var (
		ID       string
		Title    string
		AuthorID AuthorID
	)
	if err := row.Scan(&ID, &Title, &AuthorID); err != nil {
		return nil, fmt.Errorf("GetBook: %w", err)
	}

	return &Book{
		ID:    ID,
		Title: Title,
		Author: Author{
			AuthorID: AuthorID,
			name:     "test",
		},
	}, nil
}

func GetAuthorName2(ctx context.Context, title string) (string, error) {
	r := &Repository2{
		db: ch03.Db,
	}
	b, err := r.GetBook(ctx, title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("GetAuthorName2: unknown book %v", err)
		}
	}
	return b.Name(), nil
}
