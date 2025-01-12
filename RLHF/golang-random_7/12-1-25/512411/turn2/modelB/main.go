package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func makeAPICall(wg *sync.WaitGroup, apiURL string) {
	defer wg.Done()
	fmt.Printf("Making API call to: %s\n", apiURL)
	// Simulate the API call taking some time
	time.Sleep(2 * time.Second)
	// Make the actual HTTP request to the API here
	_, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error making API call to %s: %v\n", apiURL, err)
	}
}

func main() {
	var wg sync.WaitGroup

	// Define the URLs of the APIs you want to call
	apiURLs := []string{
		"https://example.com/api1",
		"https://example.com/api2",
		"https://example.com/api3",
		"https://example.com/api4",
	}

	wg.Add(len(apiURLs))

	// Start concurrent API calls
	for _, apiURL := range apiURLs {
		go makeAPICall(&wg, apiURL)
	}

	// Wait for all API calls to complete
	wg.Wait()

	fmt.Println("All API calls completed.")
}
