package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Simulate a handler function that performs some work
func handleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500) // Simulate work time
	fmt.Fprintf(w, "Hello, concurrent world!\n")
}

func main() {
	const numHandlers = 100
	var wg sync.WaitGroup

	// Start the server
	fmt.Println("Server started...")
	http.HandleFunc("/", handleRequest)

	// Run handlers concurrently
	for i := 0; i < numHandlers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			http.ListenAndServe(":8080", nil)
		}()
	}

	wg.Wait()
}
