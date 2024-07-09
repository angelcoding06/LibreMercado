package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBHost   string
	DBUser   string
	DBPass   string
	DBName   string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		Port:   getEnv("PORT", "3000"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASSWORD", "password"),
		DBName: getEnv("DB_NAME", "mydatabase"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
