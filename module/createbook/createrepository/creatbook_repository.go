package createrepository

import (
	"database/sql"
	"mylibary/module/entities"

	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type CreateBookRepository interface {
	CreateBookRepository(c *fiber.Ctx) error
}

type createBookSqlRepo struct {
	db *sql.DB
}

func GetDbCreateBook(db *sql.DB) *createBookSqlRepo {
	return &createBookSqlRepo{db: db}
}

var books []entities.Books

func (database *createBookSqlRepo) CreateBookRepository(c *fiber.Ctx) error {

	book := new(entities.Books)
	// book.Date_Write.AddDate(2021, 2, 25)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	books = append(books, *book)
	fmt.Printf("book.Book_Id: %v\n", book.Book_id)
	fmt.Printf("book.Book_Name: %v\n", book.Book_name)
	fmt.Printf("book.Book_Author: %v\n", book.Book_author)
	fmt.Printf("book.Adminupdate_Id: %v\n", book.Adminupdate_id)

	_, err := database.db.Exec("INSERT INTO public.mybook_db(book_name, book_author, adminupdate_id) "+
		"VALUES ($1, $2, $3)", book.Book_name, book.Book_author, book.Adminupdate_id)

	if err != nil {
		return err
	}

	return nil
}
