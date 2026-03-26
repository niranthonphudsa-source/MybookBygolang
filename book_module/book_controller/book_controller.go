package bookcontroller

import (
	"database/sql"
	"fmt"
	bookusecase "mylibary/book_module/book_usecase"
	"mylibary/book_module/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type bookController struct {
	usecase bookusecase.BookUsecase
}

func NewBookController(usecase bookusecase.BookUsecase) *bookController {
	return &bookController{usecase: usecase}
}

func (u *bookController) GetBookAllController(c *fiber.Ctx) error {
	books, err := u.usecase.GetBookAllRepoImpl()
	if err != nil {
		return err
	}
	return c.JSON(books)
}

func (u *bookController) GetBookIdController(c *fiber.Ctx) error {
	book_id, _ := strconv.Atoi(c.Params("book_id"))
	fmt.Print(book_id)
	books, err := u.usecase.GetBookIdRepoImpl(book_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Print(book_id)
			return c.Status(fiber.ErrBadRequest.Code).SendString("ERR No ROW")
		}
		return err
	}
	return c.JSON(books)
}

func (u *bookController) CreateBookController(c *fiber.Ctx) error {
	var bookNew = new(entities.Books)

	if err := c.BodyParser(bookNew); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error BodyParser CreateBook Controller")
	}
	fmt.Print("Controller", bookNew.Book_name, bookNew.Book_author, bookNew.Adminupdate_id)
	err := u.usecase.CreateBookRepoImpl(bookNew)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Error is: " + err.Error(),
		})
	}

	return c.JSON(bookNew)
}

func (u *bookController) DeleteBookIdController(c *fiber.Ctx) error {
	book_id, _ := strconv.Atoi(c.Params("book_id"))
	fmt.Print(book_id)
	err := u.usecase.DeleteBookIdImpl(book_id)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Delete Faild Error is " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete Success",
	})
}

func (u *bookController) UpdateBookIdController(c *fiber.Ctx) error {
	var book_update = new(entities.Books)

	book_id, _ := strconv.Atoi(c.Params("book_id"))
	fmt.Print("Controller", book_id)

	if err := c.BodyParser(&book_update); err != nil {
		return err
	}
	_, err := u.usecase.UpdateBookIdRepoImpl(book_id, book_update)

	if err != nil {
		return err
	}

	return c.JSON(&book_update)

}
