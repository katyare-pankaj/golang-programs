package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func worker(ctx context.Context, id int, message string) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d: Cancelled\n", id)
			return
		default:
			fmt.Printf("Worker %d: %s\n", id, message)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Create a parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start some workers
	for i := 1; i <= 5; i++ {
		go worker(ctx, i, fmt.Sprintf("Worker %d is running", i))
	}

	// Simulate some work
	time.Sleep(2 * time.Second)

	// Cancel the context to stop all workers
	log.Println("Cancelling workers...")
	cancel()

	// Wait for workers to exit
	time.Sleep(time.Second)
}
