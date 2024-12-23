package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a function that processes a single batch
func processBatch(batch int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this batch is completed
	fmt.Printf("Starting to process batch %d...\n", batch)
	time.Sleep(2 * time.Second) // Simulate processing work taking 2 seconds
	fmt.Printf("Batch %d processing completed.\n", batch)
}

func main() {
	// Define a list of batches
	batches := []int{1, 2, 3, 4, 5}

	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Increment the WaitGroup counter for each batch
	for _, batch := range batches {
		wg.Add(1)

		// Launch a goroutine for each batch
		go func(batch int) {
			processBatch(batch, &wg) // Pass the batch number and WaitGroup to the goroutine
		}(batch)
	}

	// Wait for all batches to be processed
	wg.Wait()

	fmt.Println("All batches have been processed.")
}
