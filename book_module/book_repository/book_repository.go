package bookrepository

import (
	"database/sql"
	"fmt"
	"mylibary/book_module/entities"

	_ "github.com/lib/pq"
)

type BooksRepository interface {
	GetBooksAll() ([]entities.Books, error)
	GetBookId(book_id int) (*entities.Books, error)
}

type getDbBooks struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *getDbBooks {
	fmt.Print(&db)
	return &getDbBooks{db: db}
}

var books []entities.Books

func (conn *getDbBooks) GetBooksAll() ([]entities.Books, error) {
	rows, err := conn.db.Query("SELECT book_id, book_name," +
		"book_autho, admibupdate_id FROM public.mybook_db")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b entities.Books
		err := rows.Scan(&b.Book_id, &b.Book_name, &b.Book_author, b.Adminupdate_id)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Print(books)
	return books, nil
}

func (conn *getDbBooks) GetBookId(book_id int) (*entities.Books, error) {
	var b entities.Books
	err := conn.db.QueryRow("SELECT book_id, book_name,"+
		"book_author, admibupdate_id FROM public.mybook"+
		"_db WHERE book_id=$1", book_id).Scan(b.Book_id, b.Book_name, b.Book_author, b.Adminupdate_id)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
