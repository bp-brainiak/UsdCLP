package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetValue(key string) string {
	godotenv.Load()
	err := godotenv.Load("prod_env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
