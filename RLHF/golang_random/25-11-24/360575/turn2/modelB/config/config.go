package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL   string
	Port          int
	EnableLogging bool
	LogLevel      string
}

const (
	defaultPort     = 8080
	defaultLogLevel = "info"
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Read DATABASE_URL from environment variable
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL environment variable not set")
	}
	cfg.DatabaseURL = databaseURL

	// Read PORT from environment variable and fall back to default
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(defaultPort)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid PORT value: %v", portStr)
	}
	cfg.Port = port

	// Read ENABLE_LOGGING from environment variable and default to false
	enableLoggingStr := os.Getenv("ENABLE_LOGGING")
	enableLogging, err := strconv.ParseBool(enableLoggingStr)
	if err != nil {
		enableLogging = false
	}
	cfg.EnableLogging = enableLogging

	// Read LOG_LEVEL from environment variable and fall back to default
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = defaultLogLevel
	}
	cfg.LogLevel = logLevel

	return cfg, nil
}
