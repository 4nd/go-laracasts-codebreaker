package config

import (
	"code-breaker/views"
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	loadEnv()

	err := views.LoadViews()
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
}
