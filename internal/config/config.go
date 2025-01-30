package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
	Port       string
	CORSOrigins []string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	corsOrigins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")

	config := &Config{
		DBHost:     MustGetEnv("DB_HOST"),
		DBUser:     MustGetEnv("DB_USER"),
		DBPassword: MustGetEnv("DB_PASSWORD"),
		DBName:     MustGetEnv("DB_NAME"),
		DBPort:     MustGetEnv("DB_PORT"),
		JWTSecret:  MustGetEnv("JWT_SECRET"),
		Port:       MustGetEnv("PORT"),
		CORSOrigins: corsOrigins,
	}
	
	return config
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
