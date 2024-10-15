package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func gzipHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the client supports gzip compression
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		http.Error(w, "Client does not support gzip compression", http.StatusNotAcceptable)
		return
	}

	// Set the content-encoding header to indicate gzip compression
	w.Header().Set("Content-Encoding", "gzip")

	// Create a gzip writer
	gz := gzip.NewWriter(w)
	defer gz.Close()

	// Wrap the response writer to write compressed data
	gzipWriter := io.Writer(gz)

	// Your original response handling logic here
	// For demonstration, we'll just write a sample response
	fmt.Fprintf(gzipWriter, "Hello, this is a compressed response!")
}

func main() {
	http.HandleFunc("/", gzipHandler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
