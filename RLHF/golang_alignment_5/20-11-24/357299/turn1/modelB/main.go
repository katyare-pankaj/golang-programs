package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func processRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Limit the number of concurrent goroutines using a channel
	sem := make(chan struct{}, 10) // Adjust the buffer size as needed

	// Process each request item in a goroutine
	for _, item := range []string{"item1", "item2", "item3"} {
		// Acquire a semaphore to limit the number of concurrent goroutines
		sem <- struct{}{}

		// Start a goroutine to process the item
		wg.Add(1)
		go func(item string) {
			defer func() {
				// Release the semaphore after the goroutine finishes
				<-sem
				// Signal that the goroutine has finished
				wg.Done()
			}()

			// Perform some work on the item
			time.Sleep(time.Second)
			fmt.Println("Processed item:", item)

			// Check for context cancellation
			select {
			case <-ctx.Done():
				// Request has been canceled, return
				return
			default:
			}
		}(item)
	}

	// Wait for all goroutines to finish
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
