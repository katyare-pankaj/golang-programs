package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func someLongRunningTask(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// Context is cancelled, handle it
			log.Printf("Task %d: Cancelled\n", id)
			return
		default:
			// Simulate work with a random sleep duration
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			log.Printf("Task %d: Working...\n", id)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start a long-running task in a goroutine
	go someLongRunningTask(ctx, 1)

	time.Sleep(3 * time.Second) // Wait for a little longer than the timeout

	log.Println("Main thread is completing.")
}
