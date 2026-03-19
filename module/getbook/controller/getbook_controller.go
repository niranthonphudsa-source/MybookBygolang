package controller

import (
	"mylibary/module/getbook/usecase"

	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	usecase usecase.BookUsecase
}

func NewBookController(u usecase.BookUsecase) *bookController {
	return &bookController{usecase: u}
}

func (control *bookController) GetBooksHandler(c *fiber.Ctx) error {
	books, err := control.usecase.GetBookAllRepoImpl()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(books)
}
