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
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Booting...")
	})
	fullApp, err := api.InitApp()
	if err != nil {
		log.Fatal("InitApp failed:", err)
	}
	app.Mount("/", fullApp)
	log.Println("Server listening on port", port)
	log.Fatal(app.Listen(":" + port))
}
