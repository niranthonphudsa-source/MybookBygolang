package bookrepository

import (
	"database/sql"
	"fmt"
	"mylibary/book_module/entities"

	_ "github.com/lib/pq"
)

var books []entities.Books

type BooksRepository interface {
	GetBooksAll() ([]entities.Books, error)
	GetBookId(book_id int) (*entities.Books, error)
	CreateBook(bookNew *entities.Books) error
	DeleteBookId(book_id int) error
	UpdateBookId(book_id int, bookUpdate *entities.Books) (*entities.Books, error)
}

type getDbBooks struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *getDbBooks {
	fmt.Print(&db)
	return &getDbBooks{db: db}
}

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
	err := conn.db.QueryRow("SELECT book_id, book_name, book_author, adminupdate_id "+
		" FROM public.mybook_db WHERE book_id = $1;", book_id).Scan(&b.Book_id, &b.Book_name, &b.Book_author, &b.Adminupdate_id)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (conn *getDbBooks) CreateBook(bookNew *entities.Books) error {

	fmt.Print("Repository ", bookNew.Book_name, bookNew.Book_author, bookNew.Adminupdate_id)

	_, err := conn.db.Exec("INSERT INTO public.mybook_db(book_name, book_author, adminupdate_id) "+
		" VALUES ($1, $2, $3)", bookNew.Book_name, bookNew.Book_author, bookNew.Adminupdate_id)

	if err != nil {
		fmt.Print("Error is", err)
		return err
	}
	return err

}

func (conn *getDbBooks) DeleteBookId(book_id int) error {
	fmt.Print(book_id)
	_, err := conn.db.Exec("DELETE FROM public.mybook_db WHERE book_id = $1", book_id)
	if err != nil {
		return err
	}
	return err
}

func (conn *getDbBooks) UpdateBookId(book_id int, bookUpdate *entities.Books) (*entities.Books, error) {
	fmt.Print("Repository", book_id)

	_, err := conn.db.Exec("UPDATE public.mybook_db SET "+
		"book_name = $1, book_author = $2 WHERE book_id = $3", bookUpdate.Book_name, bookUpdate.Book_author, book_id)

	if err != nil {
		return nil, err
	}
	return bookUpdate, nil

}
