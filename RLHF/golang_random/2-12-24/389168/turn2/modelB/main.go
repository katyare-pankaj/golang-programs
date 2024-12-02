package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Simulate a long-running task that can be canceled.
func longRunningTask(ctx context.Context, taskID int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d: Canceled\n", taskID)
			return
		default:
			// Simulate work with a random sleep.
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}
}

// Process HTTP requests with a timeout.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Create a context with a 5-second timeout.
	_, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Simulate a long-running operation that can be timed out.
	fmt.Fprintln(w, "Processing request...")
	time.Sleep(3 * time.Second)
	fmt.Fprintln(w, "Request processed successfully.")
}

func main() {
	// 1. Use context for cancellation:
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start multiple long-running tasks.
	numTasks := 5
	wg := sync.WaitGroup{}
	wg.Add(numTasks)
	for i := 0; i < numTasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			longRunningTask(ctx, taskID)
		}(i)
	}

	// After 2 seconds, cancel all tasks.
	time.AfterFunc(2*time.Second, cancel)

	// Wait for all tasks to finish.
	wg.Wait()

	// 2. Use timeouts for bound operations:
	server := http.Server{Addr: ":8080"}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Close the server after 10 seconds.
	time.AfterFunc(10*time.Second, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("HTTP server shutdown error: %v", err)
		}
		fmt.Println("HTTP server gracefully stopped")
	})

	// Handle HTTP requests.
	http.HandleFunc("/", handleRequest)
	log.Println("HTTP server listening on :8080")

}
