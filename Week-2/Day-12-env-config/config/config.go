package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  int
	DBDSN string
	Env   string
}

var AppConfig Config

func Load() {
	_ = godotenv.Load()

	AppConfig = Config{
		Port:  getEnvAsInt("PORT", 8000),
		DBDSN: getEnvOrFail("DB_DSN"),
		Env:   getEnv("APP_ENV", "development"),
	}
}

func getEnvOrFail(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Environment variable $s is required", key)
	}

	return val
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Invalid integer for %s: %v", key, err)
	}

	return val
}
