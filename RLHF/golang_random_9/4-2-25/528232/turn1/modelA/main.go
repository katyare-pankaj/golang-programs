package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// fetchData performs an HTTP GET request to the given URL
// and returns the response body or an error.
func fetchData(url string) (string, error) {
	// Create an HTTP client with a timeout.
	client := &http.Client{Timeout: 10 * time.Second}

	// Make a GET request to the URL.
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}

	// Ensure the response body is closed when the function exits.
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	// Check if the status code indicates a failure.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func main() {
	url := "https://api.example.com/data"

	// Call the fetchData function.
	data, err := fetchData(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Println("Response from API:", data)
}
