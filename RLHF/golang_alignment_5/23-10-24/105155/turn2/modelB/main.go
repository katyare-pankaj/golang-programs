package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// streamingHandler reads data from a file and streams it to the client.
func streamingHandler(w http.ResponseWriter, r *http.Request) {
	// Open the file to be streamed
	file, err := os.Open("example.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check if the client supports zero-copy transmission
	if _, ok := w.Write([]byte(file.Name())); ok != nil {
		// Create a temporary buffer for zero-copy copying
		buf := make([]byte, 32*1024) // Choose an appropriate buffer size

		// Perform zero-copy copying from the file to the client using the io.CopyBuffer function
		_, err = io.CopyBuffer(w, file, buf)
		if err != nil {
			log.Printf("Error streaming file: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	// Fall back to conventional copying if zero-copy is not supported
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("Error streaming file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Start the HTTP server on port 8080
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(streamingHandler)))
}
