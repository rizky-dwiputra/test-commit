package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Port    string
	Env     string

	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
	JWT_SECRET string
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	AppConfig = &Config{
		AppName:    getEnv("APP_NAME", "codelabs-api"),
		Port:       getEnv("PORT", "5000"),
		Env:        getEnv("ENV", "development"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "go_user"),
		DBPass:     getEnv("DB_PASS", "secret"),
		DBName:     getEnv("DB_NAME", "codelabs_fiber"),
		JWT_SECRET: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
