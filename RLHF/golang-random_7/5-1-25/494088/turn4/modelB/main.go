package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Exiting due to context cancellation")
			return
		default:
			fmt.Println("Worker: Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the worker goroutine
	go worker(ctx)

	// Wait for the worker goroutine to complete or timeout
	<-ctx.Done()
	fmt.Println("Main: Worker has finished or timed out.")
}
