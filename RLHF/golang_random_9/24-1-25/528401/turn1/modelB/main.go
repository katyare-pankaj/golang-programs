package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	// Set the maximum number of goroutines to handle requests
	maxGoroutines = 100

	// Set the timeout duration for long-running requests
	timeoutDuration = 5 * time.Second
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Simulate a long-running task by sleeping for a random duration
	duration := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(duration)
	fmt.Fprintf(w, "Request handled after %s\n", duration)
}
func main() {
	var wg sync.WaitGroup

	// Create a channel to limit the number of goroutines
	semaphore := make(chan struct{}, maxGoroutines)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Acquire a token from the semaphore to limit the number of goroutines
		semaphore <- struct{}{}
		defer func() {
			// Release the token after the goroutine exits
			<-semaphore
		}()

		// Start a new goroutine to handle the request
		wg.Add(1)
		go func() {
			defer wg.Done()
			handleRequest(w, r)
		}()

		// Wait for the goroutine to finish within the specified timeout
		c := make(chan struct{})
		go func() {
			wg.Wait()
			close(c)
		}()
		select {
		case <-c:
			// Goroutine finished within the timeout
		case <-time.After(timeoutDuration):
			// Goroutine timed out
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
		}
	})

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
