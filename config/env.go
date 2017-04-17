package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env(item string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(item)
}
