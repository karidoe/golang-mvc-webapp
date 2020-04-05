package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Unable to load .env")
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}