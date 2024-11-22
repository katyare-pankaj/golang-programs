package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Create a group to wait for all goroutines
	var wg sync.WaitGroup

	// Create a context with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start three goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Wait for all goroutines to finish or timeout
	wg.Wait()

	log.Println("Main function exiting")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work with a randomized sleep duration
	sleepDuration := time.Duration(rand.Intn(5)) * time.Second

	log.Printf("Worker %d: Starting work (will sleep for %v)\n", id, sleepDuration)

	select {
	case <-ctx.Done():
		// Context was canceled or timed out
		log.Printf("Worker %d: Cancelled\n", id)
		return
	case <-time.After(sleepDuration):
		// Work completed normally
		log.Printf("Worker %d: Work completed\n", id)
	}
}
