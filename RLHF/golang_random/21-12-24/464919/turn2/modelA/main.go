package main

import (
	"context"
	"log"
	"time"
)

// Worker simulates a worker goroutine that performs a task.
func Worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker canceled by context")
			return
		case <-time.After(time.Second):
			log.Println("Worker working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Start the worker goroutine
	go Worker(ctx)

	// Simulate a condition where we need to cancel the worker
	time.Sleep(time.Second * 2)
	log.Println("Canceling worker")
	cancel() // Cancel the context
}
