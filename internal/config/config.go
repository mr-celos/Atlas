package config

import (
	"os"
)

type Config struct {
	Port       string
	Version    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func getEnvOrDefault(key, defaultValue string) string { // returns the value of an environment variable or a default value if not set
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Load() Config { // returns a config struct with values from environment variables or defaults
	port := getEnvOrDefault("PORT", ":8080")
	version := getEnvOrDefault("PRODUCT_VERSION", "Development")
	dbhost := getEnvOrDefault("DATABASE_HOST", "localhost")
	dbport := getEnvOrDefault("DATABASE_PORT", "5432")
	dbuser := getEnvOrDefault("DATABASE_USER", "postgres")
	dbpassword := getEnvOrDefault("DATABASE_PASSWORD", "")
	dbname := getEnvOrDefault("DATABASE_NAME", "atlas")

	return Config{
		Port:       port,
		Version:    version,
		DBHost:     dbhost,
		DBPort:     dbport,
		DBUser:     dbuser,
		DBPassword: dbpassword,
		DBName:     dbname,
	}
}
