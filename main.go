package main

import (
	"log"
	"os"

	"github.com/Wenell09/MyStock/internal/api"
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
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on port", port)
	log.Fatal(app.Listen(":" + port))
}
