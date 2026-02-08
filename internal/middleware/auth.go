package middleware

import (
	"strings"

	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).
				JSON(utils.NewResponseError(fiber.StatusUnauthorized, "Unauthorized", "Token Empty!"))
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		jwkSet, err := GetJWKSet()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(utils.NewResponseError(500, "Error", err.Error()))
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			kidVal, ok := t.Header["kid"]
			if !ok {
				return nil, fiber.ErrUnauthorized
			}
			kid, ok := kidVal.(string)
			if !ok {
				return nil, fiber.ErrUnauthorized
			}
			key, ok := jwkSet.LookupKeyID(kid)
			if !ok {
				return nil, fiber.ErrUnauthorized
			}
			var raw interface{}
			if err := key.Raw(&raw); err != nil {
				return nil, err
			}
			return raw, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).
				JSON(utils.NewResponseError(fiber.StatusUnauthorized, "Unauthorized", "Token Invalid!"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).
				JSON(utils.NewResponseError(fiber.StatusUnauthorized, "Unauthorized", "Invalid token claims"))
		}
		c.Locals("user_id", claims["sub"])
		c.Locals("email", claims["email"])
		return c.Next()
	}
}
