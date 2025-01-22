package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabasePort string
	Port         string
	Env          string
	Host         string
}

var (
	config *Config
	once   sync.Once
)

// LoadConfig initializes the configuration only once
func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file if it exists
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}

		config = &Config{
			Host:         getEnv("HOST", "localhost"),
			DatabaseURL:  getEnv("DATABASE_URL", "localhost"),
			DatabasePort: getEnv("DATABASE_PORT", "3306"),
			Port:         getEnv("PORT", "8080"),
			Env:          getEnv("ENV", "development"),
		}
	})

	return config
}

// getEnv fetches a key from the environment or uses a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
