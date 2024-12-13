package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// FetchURL simulates fetching a URL and respects the context for cancellation
func FetchURL(ctx context.Context, url string) (string, error) {
	// Create a new HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	// Simulate a network operation with an HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return fmt.Sprintf("Fetched %s with status %s", url, resp.Status), nil
}

func main() {
	// Set a timeout for the whole operation (e.g., 2 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure that resources are cleaned up

	url := "http://example.com" // Example URL

	// Call the FetchURL function with the context
	resultCh := make(chan string)
	errCh := make(chan error)

	go func() {
		result, err := FetchURL(ctx, url)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case err := <-errCh:
		fmt.Printf("Error fetching URL: %v\n", err)
	case <-ctx.Done():
		// Handle the case where the context times out
		fmt.Println("Request timed out:", ctx.Err())
	}
}
