package main

import (
	"fmt"
	"log"

	"go-programs/RLHF/golang_random/29-11-24/389128/turn3/modelA/config"
)

func main() {
	// Load configuration file
	configFile := "config.json"
	config, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Check feature flags
	if config.Features.NewFeature {
		fmt.Println("New feature is enabled.")
		newFeatureDemo()
	} else {
		fmt.Println("New feature is disabled.")
	}

	if config.Features.DebugMode {
		fmt.Println("Debug mode is enabled.")
	} else {
		fmt.Println("Debug mode is disabled.")
	}
}

// newFeatureDemo performs operations specific to the new feature.
func newFeatureDemo() {
	fmt.Println("Running the new feature functionality.")
}
