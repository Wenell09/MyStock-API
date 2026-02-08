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
			return c.Status(401).JSON(utils.NewResponseError(401, "Unauthorized", "Token missing"))
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		jwkSet, err := GetJWKSet()
		if err != nil || jwkSet == nil {
			log.Println("[AUTH] JWKS error:", err)
			return c.Status(500).JSON(utils.NewResponseError(500, "Error", "Auth service unavailable"))
		}
		claims := jwt.MapClaims{}
		parser := jwt.NewParser(jwt.WithValidMethods([]string{"RS256"}))
		token, err := parser.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if t == nil || t.Header == nil {
				return nil, fiber.ErrUnauthorized
			}
			kidVal, ok := t.Header["kid"]
			if !ok || kidVal == nil {
				return nil, fiber.ErrUnauthorized
			}
			kid, ok := kidVal.(string)
			if !ok {
				return nil, fiber.ErrUnauthorized
			}
			key, ok := jwkSet.LookupKeyID(kid)
			if !ok || key == nil {
				return nil, fiber.ErrUnauthorized
			}
			var raw interface{}
			if err := key.Raw(&raw); err != nil || raw == nil {
				return nil, fiber.ErrUnauthorized
			}
			return raw, nil
		})
		if err != nil || token == nil || !token.Valid {
			return c.Status(401).JSON(utils.NewResponseError(401, "Unauthorized", "Invalid token"))
		}
		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			return c.Status(401).JSON(utils.NewResponseError(401, "Unauthorized", "Invalid token subject"))
		}
		email := ""
		if v, ok := claims["email"]; ok && v != nil {
			email, _ = v.(string)
		}
		c.Locals("user_id", sub)
		c.Locals("email", email)
		return c.Next()
	}
}
