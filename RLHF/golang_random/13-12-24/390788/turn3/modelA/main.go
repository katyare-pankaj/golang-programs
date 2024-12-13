package main

import (
	"fmt"
	"sync"
)

func main() {
	// Example array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Number of Goroutines
	numWorkers := 4
	// Create a wait group
	var wg sync.WaitGroup
	// Mutex for synchronizing access to the sum
	var mu sync.Mutex
	// Variable to hold the total sum
	totalSum := 0

	// Calculate the size of each chunk
	chunkSize := (len(arr) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Add to the wait group

		go func(start int) {
			defer wg.Done() // Signal that this Goroutine is done
			end := start + chunkSize
			if end > len(arr) {
				end = len(arr)
			}
			// Calculate partial sum
			partialSum := 0
			for j := start; j < end; j++ {
				partialSum += arr[j]
			}

			// Lock before updating the shared resource
			mu.Lock()
			totalSum += partialSum
			mu.Unlock()
		}(i * chunkSize)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Print the total sum
	fmt.Printf("Total Sum: %d\n", totalSum)
}
