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

func (c *bookController) GetBooksHandler(ctx *fiber.Ctx) error {
	books, err := c.usecase.GetBookAllRepoImpl()
	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}
	return ctx.JSON(books)
}
