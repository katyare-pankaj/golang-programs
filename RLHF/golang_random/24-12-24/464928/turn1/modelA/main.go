package main

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Database string `json:"database"`
	Port     int    `json:"port"`
}

// Example function to load configuration from a file or a network source.
func loadConfiguration(configChan chan<- Config, callback func(Config, error)) {
	// Simulate loading a configuration from an external source.
	configFile := "config.json"

	// For simplicity, assume this function reads the configuration from the configFile.
	// In a real-world scenario, this could involve reading from a database, environment
	// variables, or a remote API.
	config, err := loadConfigFromFile(configFile)
	if err != nil {
		callback(Config{}, err)
		return
	}

	configChan <- config
	callback(config, nil)
}

func loadConfigFromFile(filePath string) (Config, error) {
	// Simulate a long-running file load operation with a delay
	time.Sleep(2 * time.Second)
	return Config{Database: "testdb", Port: 8080}, nil
}

func main() {
	// Channel to receive the loaded configuration
	configChan := make(chan Config)

	// Function to be called when configuration is loaded
	configCallback := func(config Config, err error) {
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		fmt.Println("Configuration loaded:", config)
		// Use the loaded configuration in your application

		// Simulate starting a server with the loaded configuration
		go runServer(config)
	}

	// Go routine to load configuration
	go loadConfiguration(configChan, configCallback)

	// Simulate main application processing
	fmt.Println("Main application processing...")
	time.Sleep(3 * time.Second) // Simulate some work

	fmt.Println("Main application done.")
}

func runServer(config Config) {
	fmt.Println("Starting server on port", config.Port)
	http.ListenAndServe(":"+fmt.Sprintf("%d", config.Port), nil)
}
