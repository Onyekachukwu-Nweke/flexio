package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName     string `envconfig:"DB_NAME" default:"postgres"`
	DBSSLMode  string `envconfig:"DB_SSL_MODE" default:"disable"`
}

// LoadConfig loads the configuration from environment variables
func LoadConfig(envFile string) (*Config, error) {
	// Load .env file if it exists
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USERNAME", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "intrazine"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	return config, nil
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvRequired retrieves a required environment variable
// Returns empty string if the variable is not set
func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Warning: Required environment variable %s is not set\n", key)
	}
	return value
}
