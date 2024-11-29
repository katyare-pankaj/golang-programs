package feature

import (
	"fmt"
	"os"
	"strconv"
)

// GetFlag retrieves the value of a feature flag from the environment variables.
// It returns true if the flag is enabled (value is "true" or "1"), false otherwise.
func GetFlag(name string) bool {
	value := os.Getenv(name)
	if value == "" {
		return false // Default to false if the flag is not set
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Errorf("invalid flag value for %s: %v", name, err))
	}
	return b
}
