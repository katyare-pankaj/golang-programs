package main

import (
	"fmt"
	"sync"
)

// calculateSegmentSum calculates the sum of a segment of the slice.
func calculateSegmentSum(data []int, start, end int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done() // Decrement the WaitGroup counter when done.

	sum := 0
	for i := start; i < end; i++ {
		sum += data[i]
	}

	resultChan <- sum // Send the result to the result channel.
}

func main() {
	// Create a large slice of integers.
	data := make([]int, 1000)
	for i := 0; i < len(data); i++ {
		data[i] = i + 1 // Fill the slice with numbers 1 to 1000.
	}

	const numGoroutines = 5 // Number of goroutines to use for concurrent processing.

	var wg sync.WaitGroup                       // Create a WaitGroup to manage the goroutines.
	resultChan := make(chan int, numGoroutines) // Channel to collect results from goroutines.

	// Calculate the size of each segment.
	chunkSize := (len(data) + numGoroutines - 1) / numGoroutines

	// Start the goroutines to calculate segment sums.
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data) // Ensure the end index does not exceed the slice length.
		}

		wg.Add(1) // Increment the WaitGroup counter.
		go calculateSegmentSum(data, start, end, &wg, resultChan)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(resultChan) // Close the channel since no more values will be sent.

	// Collect the results and calculate the total sum.
	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Println("Total sum of elements:", totalSum)
}
