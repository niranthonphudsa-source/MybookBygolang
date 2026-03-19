package createcontroller

import (
	"mylibary/module/createbook/createusecase"
	"mylibary/module/entities"

	"github.com/gofiber/fiber/v2"
)

type createBookController struct {
	usecase createusecase.CreateBookUsecase
}

func NewCreateBookController(u createusecase.CreateBookUsecase) *createBookController {
	return &createBookController{usecase: u}
}

func (control *createBookController) CreateBooksHandler(c *fiber.Ctx) error {
	req := new(entities.Books)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, err := control.usecase.CreateBook(c)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message test": err.Error()})
	}

	return c.Status(fiber.ErrBadRequest.Code).JSON(req)
}
