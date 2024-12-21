package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Process performs a task that takes some time and can be canceled.
func Process(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the goroutine is done

	log.Println("Starting process...")

	select {
	case <-ctx.Done():
		log.Println("Process canceled by context")
		return
	case <-time.After(time.Second * 5):
		log.Println("Process completed normally")
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start 3 goroutines with a context and a waitgroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment the waitgroup counter
		go Process(ctx, &wg)
	}

	log.Println("Main routine is waiting for goroutines to finish...")

	// Simulate a condition where we need to cancel all the goroutines
	time.Sleep(time.Second * 2)
	log.Println("Canceling all processes")
	cancel() // Cancel the context

	// Wait for all goroutines to complete or be canceled
	wg.Wait()

	log.Println("All goroutines have finished.")
}
