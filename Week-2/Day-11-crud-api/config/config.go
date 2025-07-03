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
}

func Load() Config {
	_ = godotenv.Load()

	portStr := getEnv("PORT", "8000")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT: %v", err)
	}

	db := getEnv("DB_DSN", "")
	if db == "" {
		log.Fatalf("DB_DSN is required")
	}

	return Config{
		Port:  port,
		DBDSN: db,
	}
}

func getEnv(k, fallback string) string {
	v := os.Getenv(k)

	if v == "" {
		return fallback
	}

	return v
}
