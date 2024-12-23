package config

import (
	"log"
	"os"
)

var (
	DatabaseURL string
	JWTSecret   string
)

func LoadConfig() {
	DatabaseURL = os.Getenv("DATABASE_URL")
	JWTSecret = os.Getenv("JWT_SECRET")

	if DatabaseURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set in the environment variables")
	}
}