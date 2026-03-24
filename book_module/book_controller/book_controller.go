package bookcontroller

import (
	"database/sql"
	"fmt"
	bookusecase "mylibary/book_module/book_usecase"
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
