package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// simulateRequest simulates processing a request that may take a while.
// It will respect context cancellation when the request exceeds the timeout.
func simulateRequest(ctx context.Context, id int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	select {
	case <-time.After(time.Duration(1+id) * time.Second): // Simulate varying processing time
		ch <- fmt.Sprintf("Request %d processed", id)
	case <-ctx.Done():
		ch <- fmt.Sprintf("Request %d cancelled due to timeout", id)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	const timeout = 3 * time.Second
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make(chan string, 5) // Buffer for 5 requests

	// Simulate processing 5 requests concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go simulateRequest(ctx, i, &wg, results)
	}

	// Wait for all requests to complete or be cancelled
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Fprintln(w, result)
	}
}

func main() {
	http.HandleFunc("/simulate", apiHandler)
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
