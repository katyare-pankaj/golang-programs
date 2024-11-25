package envconfig

import (
	"fmt"
	"os"
)

// Config holds the configuration values read from environment variables.
type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	Port        int    `env:"PORT,required"`
	Debug       bool   `env:"DEBUG,default=false"`
}

// NewConfig creates a new Config instance by reading values from environment variables.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Read and validate required environment variables
	if err := cfg.readAndValidate("DATABASE_URL"); err != nil {
		return nil, err
	}

	if err := cfg.readAndValidate("PORT"); err != nil {
		return nil, err
	}

	// Read optional environment variable with default value
	if err := cfg.readString("DEBUG", "false"); err != nil {
		return nil, err
	}

	return cfg, nil
}

// readAndValidate reads an environment variable and ensures it is not empty.
func (cfg *Config) readAndValidate(envKey string) error {
	value := os.Getenv(envKey)
	if value == "" {
		return fmt.Errorf("environment variable %q is required", envKey)
	}
	return nil
}

// readString reads an environment variable as a string and applies a default value if missing.
func (cfg *Config) readString(envKey string, defaultValue string) error {
	value := os.Getenv(envKey)
	if value == "" {
		value = defaultValue
	}
	return nil
}
