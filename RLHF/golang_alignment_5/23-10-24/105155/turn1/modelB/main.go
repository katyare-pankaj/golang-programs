package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the HTTP server.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Assuming we have a []byte representing the content we want to send
	content := []byte("Hello, world!")

	// If the client supports zero-copy transmission, attempt to use it.
	if tcpConn, err := w.Write(content); err != nil {
		if tcpConn == 0 {
			log.Printf("Error writing response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	// If zero-copy is not supported, fall back to conventional copying.
	_, err := w.Write(content)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Start the HTTP server on port 8080
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handleRequest)))
}
