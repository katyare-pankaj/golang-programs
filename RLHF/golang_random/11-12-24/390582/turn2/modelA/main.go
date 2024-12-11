package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Simulate a long-running operation
func longRunningOperation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second): // Simulate a delay
		return nil // Operation completed successfully
	case <-ctx.Done():
		return ctx.Err() // Context was cancelled or timed out
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set a timeout for the request
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	// Perform the long-running operation
	if err := longRunningOperation(ctx); err != nil {
		if err == context.DeadlineExceeded {
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
			return
		}
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	// If successful, return a response
	fmt.Fprintln(w, "Operation completed successfully")
}

func main() {
	http.HandleFunc("/", handler)

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
