package main

import (
	"fmt"
	"sync"
)

// processSlice is a function that processes a portion of the data.
// In this example, it just multiplies each element by 2.
func processSlice(data []int, start, end int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes.

	for i := start; i < end; i++ {
		data[i] *= 2 // Example processing: multiply by 2.
	}
}

func main() {
	// Create a slice of data to be processed.
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	const numGoroutines = 3 // Number of goroutines to use.

	var wg sync.WaitGroup // Create a WaitGroup to manage the goroutines.

	// Calculate the size of each chunk of data to be processed by a goroutine.
	chunkSize := (len(data) + numGoroutines - 1) / numGoroutines

	// Start the goroutines.
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data) // Ensure the end index does not exceed the slice length.
		}

		wg.Add(1) // Increment the WaitGroup counter for each goroutine.
		go processSlice(data, start, end, &wg)
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	// Print the processed data.
	fmt.Println("Processed data:", data)
}
