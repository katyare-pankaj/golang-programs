package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func fetchData(ctx context.Context, url string) error {
	// Create a new HTTP client with the given context
	client := &http.Client{}

	// Perform the HTTP request
	resp, err := client.GetWithContext(ctx, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Simulate working with the response body
	fmt.Printf("Successfully fetched data from: %s\n", url)

	return nil
}

func main() {
	// Create a context with a deadline to simulate timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// URLs for demonstration purposes
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://example.com/does-not-exist", // This will take longer than 5 seconds
		"https://jsonplaceholder.typicode.com/comments",
	}

	// Make concurrent API calls using Goroutines
	for _, url := range urls {
		go func(u string) {
			if err := fetchData(ctx, u); err != nil {
				log.Errorf("Failed to fetch data from %s: %v\n", u, err)
			}
		}(url)
	}

	time.Sleep(6 * time.Second) // Wait for Goroutines to finish
}
