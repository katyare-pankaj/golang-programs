package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL path
		log.Printf("Method: %s, URL: %s\n", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Example handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Chain the middleware with the handler
	http.Handle("/", loggingMiddleware(handler))

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
