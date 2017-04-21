package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func Env(item string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	env_file := fmt.Sprintf("%s/.env", dir)

	err2 := godotenv.Load(env_file)

	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(item)
}
