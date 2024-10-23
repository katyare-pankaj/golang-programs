package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	bufferSize = 4096 // Adjust the buffer size as per your requirement
)

// zeroCopyReader implements io.Reader using zero-copy techniques
type zeroCopyReader struct {
	file *os.File
}

func (r *zeroCopyReader) Read(p []byte) (int, error) {
	return r.file.Read(p)
}

// ServeHTTP handles HTTP requests
func (h *zeroCopyReader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if the requested file is in the cache
	cachedFile, err := os.Open("cache/" + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cachedFile.Close()

	// Set the Content-Type header based on the file extension (you can improve this)
	contentType := "application/octet-stream"
	w.Header().Set("Content-Type", contentType)

	// Use zero-copy techniques to send the file content to the client
	stat, _ := cachedFile.Stat()
	http.ServeContent(w, r, r.URL.Path, stat.ModTime(), nil)
}

func main() {
	// Mount the cache directory at the desired path
	http.Handle("/", &zeroCopyReader{})
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
