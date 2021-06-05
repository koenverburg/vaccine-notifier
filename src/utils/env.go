package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	if os.Getenv("ENV") == "prod" {
		return os.Getenv(key)
	}

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
