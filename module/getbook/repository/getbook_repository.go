package repository

import (
	"database/sql"
	"mylibary/module/entities"

	// "github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type BookRepository interface {
	GetBookAllRepo() ([]entities.Books, error)
}

type getBookRepo struct {
	db *sql.DB
}

func GetBooksRepository(db *sql.DB) *getBookRepo {
	return &getBookRepo{db: db}
}

func (r *getBookRepo) GetBookAllRepo() ([]entities.Books, error) {

	query := `SELECT book_id, book_name, book_author, adminupdate_id FROM public.mybook_db`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Books
	for rows.Next() {
		var b entities.Books
		err := rows.Scan(&b.Book_id, &b.Book_name, &b.Book_author, &b.Adminupdate_id)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
