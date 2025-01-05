package main

import (
	"fmt"
	"net/http"
)

// LoggingMiddleware is a middleware function that logs the request method and URL path.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL path
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// HandlerFunction is a sample handler that responds to requests.
func HandlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}

func main() {
	// Create a handler that uses the logging middleware
	loggedHandler := LoggingMiddleware(http.HandlerFunc(HandlerFunction))

	// Register the handler with the HTTP server
	http.HandleFunc("/", loggedHandler)

	// Start the HTTP server
	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
