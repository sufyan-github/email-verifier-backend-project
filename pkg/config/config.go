package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct
type Config struct {
	AppPort      string
	Env          string
	SMTPTimeout  int
	LogFile      string
	AllowedHosts string
}

// Global config variable
var AppConfig *Config

// LoadConfig loads environment variables
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	timeout, err := strconv.Atoi(getEnv("SMTP_TIMEOUT", "5"))
	if err != nil {
		timeout = 5
	}

	AppConfig = &Config{
		AppPort:      getEnv("APP_PORT", "8080"),
		Env:          getEnv("ENV", "development"),
		SMTPTimeout:  timeout,
		LogFile:      getEnv("LOG_FILE", "app.log"),
		AllowedHosts: getEnv("ALLOWED_HOSTS", "*"),
	}
}

// Helper function
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}