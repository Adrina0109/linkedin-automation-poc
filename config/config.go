package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    string
	Password string
	DryRun   bool
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Email:    os.Getenv("LOGIN_EMAIL"),
		Password: os.Getenv("LOGIN_PASSWORD"),
		DryRun:   os.Getenv("DRY_RUN") == "true",
	}
}
