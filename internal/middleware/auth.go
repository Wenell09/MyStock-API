package middleware

import (
	"strings"

	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.NewResponseError(fiber.StatusUnauthorized, "Unauthorized", "Token Empty!"))
		}
		jwkSet, err := GetJWKSet()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NewResponseError(fiber.StatusInternalServerError, "Unauthorized", "Error Get JWK!"))
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			kid := t.Header["kid"].(string)
			key, ok := jwkSet.LookupKeyID(kid)
			if !ok {
				return nil, fiber.ErrUnauthorized
			}
			var raw interface{}
			key.Raw(&raw)
			return raw, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.NewResponseError(fiber.StatusUnauthorized, "Unauthorized", "Token Invalid!"))
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["sub"])
		c.Locals("email", claims["email"])
		return c.Next()
	}
}
