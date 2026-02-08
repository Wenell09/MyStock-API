package main

import (
	"log"
	"os"

	"github.com/Wenell09/MyStock/internal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app, err := api.InitApp()
	if err != nil || app == nil {
		log.Println("InitApp failed, starting minimal fiber")
		app = fiber.New()
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Booting...")
		})
	}
	log.Println("Listening on", port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
