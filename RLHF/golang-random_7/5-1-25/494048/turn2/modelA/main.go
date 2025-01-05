package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// LogMiddleware is a middleware function that logs the request processing time.
func LogMiddleware(next http.Handler) http.Handler {
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
	// Example handler that simulates some processing time
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Simulate random processing time
		w.Write([]byte("Hello, World!"))
	})

	// Chain the middleware with the handler
	http.Handle("/", LogMiddleware(handler))

	// Start the HTTP server
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
