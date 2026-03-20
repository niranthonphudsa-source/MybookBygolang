package controller

import (
	"database/sql"
	"fmt"
	"mylibary/module/getbook/usecase"

	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	usecase usecase.BookUsecase
}

type bookIdController struct {
	usecase1 usecase.BookIdUsecase
}

func NewBookController(u usecase.BookUsecase) *bookController {
	return &bookController{usecase: u}
}

func NewBookIdController(u usecase.BookIdUsecase) *bookIdController {
	return &bookIdController{usecase1: u}
}

func (control *bookController) GetBooksHandler(c *fiber.Ctx) error {
	books, err := control.usecase.GetBookAllRepoImpl()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(books)
}

func (control *bookIdController) GetBookIdHandler(c *fiber.Ctx) error {
	book_id := c.Params("book_id")
	books, err := control.usecase1.GetBookIdRepoImpl(book_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Print(book_id)
			return c.Status(fiber.ErrBadRequest.Code).SendString("ERR No ROW")
		}
		return err
	}
	return c.JSON(books)
}
