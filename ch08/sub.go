package ch08

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"webapp-dev-in-golang/ch03"

	"github.com/go-sql-driver/mysql"
)

// func init() {
// 	connectToPostgres()
// 	connectToMySQL()
// }

func Sub() {
	// list_8_3()
	// list_8_6()
	// list_8_9()
	// list_8_14()
	list_8_15()
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
		db: ch03.PgDb,
	}
	b, err := r.GetBook(ctx, title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("GetAuthorName2: unknown book %v", err)
		}
	}
	return b.Name(), nil
}

const MySQLDuplicateEntryErrorCode uint16 = 1062

var errAlreadyExists = errors.New("deplicate entry")

func list_8_9() {
	r := NewRepo()
	b := &Book{
		ID:    "0000001",
		Title: "webapp development in golang",
		Author: Author{
			AuthorID: "0001",
			name:     "Shinizu-san",
		},
	}
	err := r.SaveBook(context.Background(), b)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Repo) SaveBook(ctx context.Context, book *Book) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO books(id, title, author_id) VALUE(?, ?, ?)", book.ID, book.Title, book.AuthorID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr.Number == MySQLDuplicateEntryErrorCode {
			return fmt.Errorf("store: cannot save book_id %s: %w", book.ID, errAlreadyExists)
		}
		return fmt.Errorf("SaveBook: %w", err)
	}
	return nil
}
