package main

import (
	"fmt"
	"net/http"
	"time"
)

// Constants for rate limiting
const (
	maxTokens       = 5               // Maximum number of tokens in the bucket
	tokenRefillRate = 1 * time.Second // Rate at which tokens are added (1 token per second)
)

// Token bucket implementation
var tokenBucket = make(chan struct{}, maxTokens)

// Refill the token bucket at the specified rate
func refillTokenBucket() {
	for {
		tokenBucket <- struct{}{}
		time.Sleep(tokenRefillRate)
	}
}

// Rate limiting middleware
func rateLimitMiddleware(next http.Handler) http.Handler {
	go refillTokenBucket() // Start refilling the token bucket
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Wait for a token to become available
		<-tokenBucket
		defer func() { tokenBucket <- struct{}{} }() // Release the token after the request

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

// Main HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request served at", time.Now())
}

func main() {
	// Apply rate limiting middleware to the main handler
	http.HandleFunc("/", rateLimitMiddleware(http.HandlerFunc(handler)))

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
