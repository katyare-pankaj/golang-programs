package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func processItem(ctx context.Context, item string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the goroutine has finished

	// Simulate some work
	time.Sleep(time.Second)

	// Check if the context has been canceled
	select {
	case <-ctx.Done():
		log.Printf("Item '%s' processing canceled\n", item)
		return
	default:
		log.Printf("Processed item: %s\n", item)
	}
}

func concurrentProcessor(ctx context.Context, items []string) {
	var wg sync.WaitGroup

	// Limit the number of concurrent goroutines
	sem := make(chan struct{}, 10)

	for _, item := range items {
		// Acquire a semaphore to limit concurrency
		sem <- struct{}{}

		// Start a goroutine to process the item
		wg.Add(1)
		go func() {
			defer func() {
				// Release the semaphore after the goroutine finishes
				<-sem
				wg.Done()
			}()
			processItem(ctx, item, &wg)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func main() {
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start concurrent processing
	log.Println("Starting concurrent processing...")
	concurrentProcessor(ctx, items)
	log.Println("Concurrent processing completed.")
}
