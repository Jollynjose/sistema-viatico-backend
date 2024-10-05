package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const DEVELOPMENT string = "dev"
const PRODUCTION string = "prod"

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	PORT        string
	ENV         string
}

func checkConfig() error {
	message := ""
	if os.Getenv("DB_HOST") == "" {
		message += "DB_HOST is not set\n"
	}
	if os.Getenv("DB_PORT") == "" {
		message += "DB_PORT is not set\n"
	}
	if os.Getenv("DB_USER") == "" {
		message += "DB_USER is not set\n"
	}
	if os.Getenv("DB_PASSWORD") == "" {
		message += "DB_PASSWORD is not set\n"
	}
	if os.Getenv("DB_NAME") == "" {
		message += "DB_NAME is not set\n"
	}
	if os.Getenv("PORT") == "" {
		message += "PORT is not set\n"
	}

	if message != "" {
		return errors.New(message)
	}
	return nil
}

func NewConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if err := checkConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	env := os.Getenv("ENV")

	if env != PRODUCTION || env == "" {
		env = DEVELOPMENT
	}

	return &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		PORT:        os.Getenv("PORT"),
		ENV:         os.Getenv("ENV"),
	}
}
