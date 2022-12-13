package main

import (
	"GO-pracitce/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.ConnectDb()

	config.Migrate()
}
