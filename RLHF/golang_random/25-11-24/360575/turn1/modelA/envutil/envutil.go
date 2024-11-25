package envutil

import "os"

// GetEnv retrieves the value of the specified environment variable.
// If the variable is not set, it returns the default value provided.
func GetEnv(key string, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
