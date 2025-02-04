package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// fetchData performs an HTTP GET request and handles response cleanup safely.
func fetchData(url string) (string, error) {
	// Create an HTTP client with a timeout.
	client := &http.Client{Timeout: 10 * time.Second}

	// Make a GET request.
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}

	// Use defer to ensure the response body is closed.
	closed := false
	defer func() {
		// Attempt to close the response body.
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Printf("error closing response body: %v", closeErr)
		} else {
			closed = true
		}
		if !closed {
			log.Println("response body was not closed properly")
		}
	}()

	// Check for any non-success status codes.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func main() {
	url := "https://api.example.com/data"

	// Fetch data from the URL.
	data, err := fetchData(url)
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}

	fmt.Println("Response from API:", data)
}
