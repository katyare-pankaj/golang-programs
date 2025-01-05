package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a WaitGroup to ensure that the middleware waits for the handler to complete
		var wg sync.WaitGroup

		// Increment the wait counter
		wg.Add(1)

		go func() {
			// Call the next handler
			defer wg.Done()
			next.ServeHTTP(w, r)
		}()

		// Wait for the handler to finish
		wg.Wait()

		// Log the duration of the request
		duration := time.Since(start)
		fmt.Printf("Request processed in %v\n", duration)
	})
}

func main() {
	// Example handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // Simulate processing time
		w.Write([]byte("Hello, World!"))
	})

	// Chain the middleware with the handler
	http.Handle("/", logMiddleware(handler))

	// Start the HTTP server
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
