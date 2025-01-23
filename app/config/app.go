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
			Host:         GetEnv("HOST", "localhost"),
			DatabaseURL:  GetEnv("DATABASE_URL", "localhost"),
			DatabasePort: GetEnv("DATABASE_PORT", "3306"),
			Port:         GetEnv("PORT", "8080"),
			Env:          GetEnv("ENV", "development"),
		}
	})

	return config
}

// getEnv fetches a key from the environment or uses a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
