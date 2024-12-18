package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func fetchData(ctx context.Context, url string) error {
	// Set timeout for the HTTP request
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err // Return error if HTTP request fails
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
	ctx := context.Background()
	// URLs for demonstration purposes
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://example.com/does-not-exist",
		"https://jsonplaceholder.typicode.com/comments",
	}

	for _, url := range urls {
		go func(u string) {
			if err := fetchData(ctx, u); err != nil {
				log.Printf("Failed to fetch data from %s: %v\n", u, err)
			}
		}(url)
	}

	time.Sleep(6 * time.Second) // Wait for Goroutines to finish
}
