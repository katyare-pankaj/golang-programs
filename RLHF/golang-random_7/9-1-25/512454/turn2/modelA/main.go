package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Calculate the factorial of a number using concurrency.
func factorialConcurrently(n int) int {
	// Initialize a wait group for the worker goroutines
	var wg sync.WaitGroup

	// Create a channel to receive the results of partial products
	resultChannel := make(chan int)

	// Calculate the segment size
	segmentSize := 1000

	// Function to calculate the partial product for a given range
	func partialProduct(start, end int) {
		var product int
		for i := start; i <= end; i++ {
			product *= i
		}
		resultChannel <- product
	}

	// Start worker goroutines for each segment
	for start := 1; start <= n; start += segmentSize {
		end := min(start+segmentSize-1, n)
		wg.Add(1)
		go func() {
			defer wg.Done()
			partialProduct(start, end)
		}()
	}

	// Aggregate the results from the worker goroutines
	var totalProduct int
	for i := 0; i < (n/segmentSize)+1; i++ {
		totalProduct *= <-resultChannel
	}

	wg.Wait()

	return totalProduct
}

// Helper function to ensure end does not exceed n
func min(end, n int) int {
	if end > n {
		return n
	}
	return end
}

func main() {
	// Example usage
	n := 100
	result := factorialConcurrently(n)
	fmt.Printf("Factorial of %d is: %d\n", n, result)
}