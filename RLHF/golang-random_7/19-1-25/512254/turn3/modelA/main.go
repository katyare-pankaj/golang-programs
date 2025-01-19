package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// makeRequest performs an HTTP GET request
func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// executeRequestWithRetry handles retries with exponential backoff
func executeRequestWithRetry(url string, maxRetries int, initialBackoff time.Duration) ([]byte, error) {
	backoff := initialBackoff

	for attempt := 0; attempt < maxRetries; attempt++ {
		log.Printf("Attempt %d for %s\n", attempt+1, url)

		resp, err := makeRequest(url)
		if err != nil {
			log.Printf("Request failed: %v\n", err)
			time.Sleep(backoff)
			backoff *= 2 // Exponential backoff
			continue
		}
		defer resp.Body.Close()

		// Read and return the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v\n", err)
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		// Return if successful
		return body, nil
	}

	return nil, fmt.Errorf("failed to fetch %s after %d attempts", url, maxRetries)
}

func main() {
	url := "http://example.com"
	maxRetries := 3
	initialBackoff := 2 * time.Second

	body, err := executeRequestWithRetry(url, maxRetries, initialBackoff)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	} else {
		fmt.Println("Response Body:", string(body))
	}
}
