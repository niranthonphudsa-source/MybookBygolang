package middleware

// import (
// 	"golang.org/x/crypto"
// 	"github.com/golang-jwt/jwt/v4"
// )

// func authRequired(c *fiber.Ctx) error {
//   cookie := c.Cookies("jwt")

//   token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//       return jwtSecretKey, nil
//   })

//   if err != nil || !token.Valid {
//       return c.SendStatus(fiber.StatusUnauthorized)
//   }

//   return c.Next()
// }
