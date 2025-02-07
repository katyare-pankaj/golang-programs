package main

import (
	"fmt"
	"sync"
)

// ConcurrentMap applies a transformation function to each element of a slice concurrently.
func ConcurrentMap[T any, U any](input []T, fn func(T) U) []U {
	numWorkers := 4 // Number of goroutines to use
	inputLen := len(input)
	output := make([]U, inputLen)
	chunkSize := (inputLen + numWorkers - 1) / numWorkers // Calculate chunk size

	var wg sync.WaitGroup
	mu := sync.Mutex{} // Mutex to protect access to the output slice

	// Worker function to process chunks
	worker := func(start, end int) {
		defer wg.Done()
		localOutput := make([]U, end-start)
		for i := start; i < end; i++ {
			localOutput[i-start] = fn(input[i])
		}
		mu.Lock()
		copy(output[start:end], localOutput)
		mu.Unlock()
	}

	// Launch workers
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > inputLen {
			end = inputLen
		}
		wg.Add(1)
		go worker(start, end)
	}

	wg.Wait() // Wait for all workers to finish
	return output
}

func main() {
	// Example usage
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Concurrently square each number
	squared := ConcurrentMap(nums, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// Verify the result
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	fmt.Println("Correctness:", equal(squared, expected))
}

// Helper function to check equality of two slices of integers
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
