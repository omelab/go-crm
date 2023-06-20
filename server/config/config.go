// config/config.go
package config

import (
	"os"
	"strconv"
)

// Application configuration
var (
	AppName      string
	AppPort      int
	DatabaseURL  string
	SecretKey    string
)


func init() {
	// Retrieve environment variables or use default values
	AppName = getEnv("APP_NAME", "YourAppName")
	AppPort = getEnvAsInt("APP_PORT", 3000)
	DatabaseURL = getEnv("DATABASE_URL", "your-postgres-connection-string")
	SecretKey = getEnv("SECRET_KEY", "your-secret-key")
}

// Helper function to retrieve string environment variable
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to retrieve integer environment variable
func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	// Handle conversion error
	port, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return port
}