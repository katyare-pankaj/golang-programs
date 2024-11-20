package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func processItem(ctx context.Context, item string) error {
	// Simulate work by sleeping for a random duration
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	select {
	case <-ctx.Done():
		// Context is canceled, return early
		return ctx.Err()
	case <-time.After(duration):
		// Perform the actual work on the item
		fmt.Println("Processed item:", item)
		return nil
	}
}

func worker(ctx context.Context, items <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range items {
		if err := processItem(ctx, item); err != nil {
			// Handle the error if processing fails
			fmt.Printf("Error processing item %s: %v\n", item, err)
		}
	}
}

func processRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Number of workers to use
	numWorkers := 5

	// Channel to send items to workers
	items := make(chan string, numWorkers)

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, items, &wg)
	}

	// Send items to the worker pool
	for _, item := range []string{"item1", "item2", "item3", "item4", "item5", "item6", "item7"} {
		items <- item
	}

	// Close the channel to signal workers that no more items will be sent
	close(items)

	// Wait for all workers to finish
	wg.Wait()

	// Write the response to the client
	fmt.Fprintln(w, "All items processed successfully")
}

func main() {
	// Create a handler for the API endpoint
	http.HandleFunc("/api/process", func(w http.ResponseWriter, r *http.Request) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Process the request
		processRequest(ctx, w, r)
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
