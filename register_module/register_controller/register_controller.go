package registercontroller

import (
	"fmt"
	"mylibary/register_module/entities"
	registerusecase "mylibary/register_module/register_usecase"

	"github.com/gofiber/fiber/v2"
)

type RegisterController struct {
	uscase registerusecase.RegisterUsecase
}

func NewRegisterController(uscase registerusecase.RegisterUsecase) *RegisterController {
	return &RegisterController{uscase: uscase}
}

func (u *RegisterController) RegisterController(c *fiber.Ctx) error {
	var newUser = new(entities.Users)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}
	err := u.uscase.RegisterUsecase(newUser)
	if err != nil {
		return err
	}
	fmt.Print("Controller", newUser)
	return c.JSON(&newUser)
}
