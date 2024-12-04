package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// fetchURL simulates a long-running operation
func fetchURL(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Simulate HTTP request
	select {
	case <-time.After(2 * time.Second): // Simulate work by sleeping
		fmt.Println("Fetched URL:", url)
		return nil
	case <-ctx.Done():
		return ctx.Err() // Propagate cancellation error
	}
}

func main() {
	// Set a timeout for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	url := "http://example.com"

	// Fetch URL with a timeout
	err := fetchURL(ctx, url)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("fetchURL timed out")
		} else if err == context.Canceled {
			fmt.Println("fetchURL was canceled")
		} else {
			fmt.Printf("fetchURL error: %v\n", err)
		}
	}
}
