package middleware

import (
	"log"
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
				JSON(utils.NewResponseError(401, "Unauthorized", "Token missing"))
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		jwkSet, err := GetJWKSet()
		if err != nil {
			log.Println("[AUTH] JWKS error:", err)
			return c.Status(500).
				JSON(utils.NewResponseError(500, "Error", "Auth service unavailable"))
		}
		parser := jwt.NewParser(
			jwt.WithValidMethods([]string{"RS256"}),
			jwt.WithAudience("authenticated"),
		)
		claims := jwt.MapClaims{}
		token, err := parser.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			kid, ok := t.Header["kid"].(string)
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
			return c.Status(401).
				JSON(utils.NewResponseError(401, "Unauthorized", "Invalid token"))
		}
		sub, ok := claims["sub"].(string)
		if !ok {
			return c.Status(401).
				JSON(utils.NewResponseError(401, "Unauthorized", "Invalid token subject"))
		}
		email, _ := claims["email"].(string)
		c.Locals("user_id", sub)
		c.Locals("email", email)

		return c.Next()
	}
}
