package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Create a new HTTP handler
	handler := http.HandlerFunc(HelloHandler)

	// Set up an HTTP server and bind the handler to the '/' route
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
