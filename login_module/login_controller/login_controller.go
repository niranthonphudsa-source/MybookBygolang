package logincontroller

import (
	"fmt"
	"mylibary/login_module/entities"
	loginusecase "mylibary/login_module/login_usecase"

	"github.com/gofiber/fiber/v2"
)

type loginController struct {
	usecase loginusecase.Loginusecase
}

func NewLoginController(usecase loginusecase.Loginusecase) *loginController {
	return &loginController{usecase: usecase}
}

func (u *loginController) LoginController(c *fiber.Ctx) error {
	var loginReq = new(entities.UserLogin)
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	fmt.Print("loginReq", loginReq)
	err := u.usecase.LoginUsecase(loginReq)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Login Failed",
		})
	}
	return c.Status(fiber.StatusOK).JSON(loginReq)
}
