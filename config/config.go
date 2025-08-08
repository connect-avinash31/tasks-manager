// package config handles the configuration for the task manager system.
package config

import "os"

type Config struct {
	DatabaseURL string `json:"database_url"` // URL for the database connection
	Port        string `json:"port"`         // Port for the task manager service
}

// LoadConfig loads the configuration from environment variables or defaults
func LoadConfig() *Config {
	return &Config{
		DatabaseURL: getEnviromentValue("DATABASE_URL", "postgres://username:password@IP:5432/task?sslmode=disable"),
		Port:        getEnviromentValue("PORT", "8080"), // Default port is 8080
	}
}

// func getEnviromentValue retrieves the value of an environment variable or returns a default value if not set.
func getEnviromentValue(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
