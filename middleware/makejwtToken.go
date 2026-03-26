package middleware

import (
	"fmt"
	"mylibary/register_module/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func extractUserFromJWT(c *fiber.Ctx) error {
	user := &entities.Users{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims{	
		"user_id": 123,
		"role":    "admin", // นี่คือต้นทางของค่า role
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	fmt.Println(claims)

	user.Email = claims["email"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}
