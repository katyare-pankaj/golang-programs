package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/chunked", chunkedHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func chunkedHandler(w http.ResponseWriter, r *http.Request) {
	// Sample data to be sent in chunks
	data := strings.Repeat("This is a test message!\n", 1000)

	// Enable chunked encoding
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	// Compress the data using Gzip
	gw := gzip.NewWriter(w)
	defer gw.Close()

	// Write the data in chunks
	chunkSize := 12 // Adjust the chunk size as needed
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		fmt.Fprintf(gw, "%x\r\n%s\r\n", len(chunk), chunk)
	}

	// Signal the end of the last chunk
	fmt.Fprintf(gw, "0\r\n\r\n")
}
