package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Example third-party function
func thirdPartyFunction(url string) error {
	log.Printf("Calling third-party function with URL: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

// Function demonstrating proper use of defer with error handling
func safeFunctionCall(url string) error {
	log.Printf("Entering safeFunctionCall...\n")

	defer log.Printf("Exiting safeFunctionCall...\n")

	err := thirdPartyFunction(url)
	if err != nil {
		log.Printf("Error in third-party function: %v\n", err)
		return err
	}

	log.Printf("Third-party function completed successfully.\n")

	// Further processing here, if any

	return nil
}

// Main function to demonstrate concurrent usage
func main() {
	url := "https://example.com"
	err := safeFunctionCall(url)
	if err != nil {
		fmt.Printf("Error in safeFunctionCall: %v\n", err)
	}

	// Optionally, you can wait for any goroutines or additional operations here
	time.Sleep(1 * time.Second)
}
