package main

import (
	"log"

	"github.com/Wenell09/MyStock/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app, err := api.InitApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listen(":8000"))
}
