package config

import (
	"os"
)

type Config struct {
	Port    string
	Version string
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
	version := getEnvOrDefault("PRODUCT_VERSION", "0.0.128934")

	return Config{
		Port:    port,
		Version: version,
	}
}
