package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Simulate a network request that can be canceled
func fetchData(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(3 * time.Second):
	}
	return "Data fetched", nil
}

// Simulate a processing task that can be canceled
func processData(ctx context.Context, data string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(2 * time.Second):
	}
	fmt.Println("Data processed:", data)
	return nil
}

func main() {
	// Create a background context
	ctx := context.Background()

	// Create a child context with a deadline for fetching data
	fetchCtx, cancelFetch := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFetch()

	// Start fetching data in a goroutine
	go func() {
		data, err := fetchData(fetchCtx)
		if err != nil {
			log.Printf("Fetch data failed: %v\n", err)
			return
		}
		log.Println("Data fetched:", data)

		// Create a child context for processing data
		processCtx, cancelProcess := context.WithTimeout(ctx, 2*time.Second)
		defer cancelProcess()

		// Start processing data in a goroutine
		go func() {
			err := processData(processCtx, data)
			if err != nil {
				log.Printf("Process data failed: %v\n", err)
			}
		}()
	}()

	// Simulate a user canceling the operation
	select {
	case <-time.After(1.5 * time.Second):
		log.Println("Canceling operation...")
		cancelFetch() // Cancel the fetch operation
	case <-time.After(5 * time.Second):
		log.Println("Operation timed out.")
	}
}
