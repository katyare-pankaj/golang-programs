package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// checkFeatureFlag reads an environment variable and checks if it is enabled.
func checkFeatureFlag(flagName string) bool {
	flagValue := os.Getenv(flagName)
	if flagValue == "" {
		return false // Feature is disabled by default
	}

	enabled, err := strconv.ParseBool(flagValue)
	if err != nil {
		log.Fatalf("Invalid value for environment variable %s: %v", flagName, err)
	}

	return enabled
}

// newFeatureDemo performs operations specific to the new feature.
func newFeatureDemo() {
	fmt.Println("Running the new feature functionality.")
}

func main() {
	if checkFeatureFlag("NEW_FEATURE_ENABLED") {
		newFeatureDemo()
	} else {
		log.Println("New feature is disabled.")
	}
}
