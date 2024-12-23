package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DBUrl string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	DBUrl = os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		panic("DATABASE_URL is not set in .env file")
	}

	fmt.Println("Configuration loaded successfully")
}