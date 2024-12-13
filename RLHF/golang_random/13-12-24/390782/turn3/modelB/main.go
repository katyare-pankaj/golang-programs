package main

import (
	"context"
	"log"
	"time"
)

// Simulates a long-running operation
func slowOperation(ctx context.Context, id int, duration time.Duration) {
	log.Printf("Goroutine %d started: sleeping for %v\n", id, duration)
	select {
	case <-time.After(duration):
		log.Printf("Goroutine %d completed\n", id)
	case <-ctx.Done():
		log.Printf("Goroutine %d cancelled\n", id)
		return
	}
}

func main() {
	// Create a background context
	ctx := context.Background()

	// Set a deadline for the operations
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()

	// Start multiple slow operations with different durations
	for i := 1; i <= 3; i++ {
		duration := time.Duration(i) * time.Second
		go slowOperation(ctx, i, duration)
	}

	// Wait for all operations to complete or timeout
	time.Sleep(3 * time.Second)

	log.Println("Main function exiting")
}
