package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// ResponseBufferPool is a sync.Pool for reusing response buffers.
var ResponseBufferPool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

// Handler reuses memory for response buffers.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Get a response buffer from the pool.
	buf := ResponseBufferPool.Get().(*strings.Builder)
	defer func() {
		// Reset the buffer and put it back in the pool for reuse.
		buf.Reset()
		ResponseBufferPool.Put(buf)
	}()

	// Generate the response content.
	buf.WriteString("Hello, World!")
	buf.WriteString(fmt.Sprintf("Requested URL: %s", r.URL.Path))

	// Write the response to the client.
	w.Write([]byte(buf.String()))
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
