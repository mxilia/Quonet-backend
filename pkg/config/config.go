package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DatabaseDSN string

	JWTSecret string

	FRONTEND_URL string
}

func GetConfig(env string) *Config {
	envFile := ".env"
	if env != "" {
		envFile = ".env." + env
	}
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	cfg := &Config{
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBName:       getEnv("DB_NAME", "test"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		JWTSecret:    getEnv("JWT_SECRET", "doodoo123"),
		FRONTEND_URL: getEnv("FRONTEND_URL", "http://localhost:3000"),
	}

	cfg.DatabaseDSN = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	return cfg
}

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
