package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from a file
func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// IsFeatureEnabled checks if a feature is enabled based on the configuration
func IsFeatureEnabled(feature string) bool {
	return viper.GetBool("features." + feature)
}

func main() {
	// Load configuration from config.yaml file
	LoadConfig("config.yaml")

	if IsFeatureEnabled("newFeature") {
		fmt.Println("New feature is enabled.")
		// Perform operations for the new feature
	} else {
		fmt.Println("New feature is disabled.")
	}

	if IsFeatureEnabled("experimentalFeature") {
		fmt.Println("Experimental feature is enabled.")
		// Perform operations for the experimental feature
	} else {
		fmt.Println("Experimental feature is disabled.")
	}
}
