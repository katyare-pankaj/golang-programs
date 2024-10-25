package main

import (
	"fmt"
	"net/http"
)

func handleLargeResponse(w http.ResponseWriter, r *http.Request) {
	const responseSize = 1024 * 1024 // 1MB response
	// Pre-allocate a buffer of the expected size
	buf := make([]byte, responseSize)

	// Generate response data (for demonstration, we'll just fill with zeros)
	for i := 0; i < responseSize; i++ {
		buf[i] = '0'
	}

	// Write the response to the client
	w.Write(buf)
}

func main() {
	http.HandleFunc("/large", handleLargeResponse)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
