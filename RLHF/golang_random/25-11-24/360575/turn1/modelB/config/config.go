package config

import (
	"fmt"
	"os"
)

// Define a struct to hold environment variable values
type Config struct {
	DatabaseURL string
}

// NewConfig creates a new Config instance by retrieving values from environment variables
func NewConfig() (*Config, error) {
	// Read the environment variable values
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Create a new Config instance and initialize it with the read values
	return &Config{DatabaseURL: databaseURL}, nil
}
