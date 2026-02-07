package main

import (
	"log"
	"os"

	"github.com/Wenell09/MyStock/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	app, err := api.InitApp()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
