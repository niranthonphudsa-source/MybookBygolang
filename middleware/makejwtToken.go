package middleware

import (
	"mylibary/middleware/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func extractUserFromJWT(c *fiber.Ctx) error {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid claims"})
	}

	user := &entities.UserCheckmiddleware{}

	if role, ok := claims["status"].(string); ok {
		user.Status = role
	}
	if email, ok := claims["email"].(string); ok {
		user.Email = email
	}

	c.Locals("currentUser", user)

	return c.Next()
}
