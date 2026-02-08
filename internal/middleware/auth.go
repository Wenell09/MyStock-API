package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Wenell09/MyStock/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type SupabaseUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func validateTokenSupabase(token string) (*SupabaseUser, error) {
	url := os.Getenv("SUPABASE_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")
	if url == "" || anonKey == "" {
		return nil, errors.New("SUPABASE_URL or SUPABASE_ANON_KEY not set")
	}
	req, _ := http.NewRequest("GET", url+"/auth/v1/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("apikey", anonKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("invalid token")
	}
	var user SupabaseUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New("invalid token: missing user id")
	}
	return &user, nil
}

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(utils.NewResponseError(401, "Unauthorized", "Token missing"))
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := validateTokenSupabase(token)
		if err != nil {
			return c.Status(401).JSON(utils.NewResponseError(401, "Unauthorized", "Invalid token"))
		}

		c.Locals("user_id", user.ID)
		c.Locals("email", user.Email)

		return c.Next()
	}
}
