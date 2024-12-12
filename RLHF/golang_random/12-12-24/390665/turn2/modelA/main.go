package main

import (
	"fmt"
	"sync"
)

// Process simulates a processing function for a part of a dataset
func processData(part int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this Goroutine is done

	// Simulate some processing work
	result := part * 2 // Example processing (e.g., doubling the number)
	fmt.Printf("Processed part %d: Result = %d\n", part, result)
}

func main() {
	var wg sync.WaitGroup

	// Simulate a dataset split into parts
	parts := []int{1, 2, 3, 4, 5}

	// Start Goroutines for each part of the dataset
	for _, part := range parts {
		wg.Add(1)                 // Increment the WaitGroup counter
		go processData(part, &wg) // Start a Goroutine for processing
	}

	// Wait for all Goroutines to finish
	wg.Wait() // This will block until all Goroutines are done

	fmt.Println("All parts processed.")
}
