package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a functional interface for callbacks
type Callback func(ctx context.Context) error

// StartWorker starts a worker goroutine that performs the specified callback with the given context.
// The worker will continue to run until the context is canceled.
func StartWorker(ctx context.Context, callback Callback, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Context is canceled, exit the worker goroutine
			return
		default:
			// Call the callback function
			if err := callback(ctx); err != nil {
				fmt.Printf("Worker error: %v\n", err)
			}
			// Add a delay before calling the callback again
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a WaitGroup to wait for all workers to complete
	var wg sync.WaitGroup

	// Start multiple worker goroutines with different callbacks
	wg.Add(2)
	go StartWorker(ctx, Worker1, &wg)
	go StartWorker(ctx, Worker2, &wg)

	// Simulate some work being done by the main goroutine
	time.Sleep(5 * time.Second)

	// Cancel the context to stop all workers
	cancel()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers stopped.")
}

// Worker1 is a callback function that performs some work.
func Worker1(ctx context.Context) error {
	fmt.Println("Worker1 is working...")
	// Simulate work by sleeping for a random duration
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return nil
}

// Worker2 is another callback function that performs some work.
func Worker2(ctx context.Context) error {
	fmt.Println("Worker2 is working...")
	// Simulate work by sleeping for a random duration
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return nil
}
