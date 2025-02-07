package main

import (
	"fmt"
	"sync"
)

// ParallelTransform applies a transformation function to each element of a slice in parallel.
func ParallelTransform[T any, U any](input []T, transform func(T) U) []U {
	numWorkers := 4 // Number of goroutines to use, this can be tuned based on the system's capabilities
	inputLen := len(input)
	output := make([]U, inputLen)

	// Channel to handle indices and transformed data
	type Job struct {
		Index int
		Value U
	}

	jobs := make(chan Job, inputLen)
	var wg sync.WaitGroup

	// Worker function to process elements
	worker := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			result := transform(input[i])
			jobs <- Job{Index: i, Value: result}
		}
	}

	// Launch workers
	chunkSize := (inputLen + numWorkers - 1) / numWorkers
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > inputLen {
			end = inputLen
		}
		wg.Add(1)
		go worker(start, end)
	}

	// Close the jobs channel after all workers are done
	go func() {
		wg.Wait()
		close(jobs)
	}()

	// Gather results
	for job := range jobs {
		output[job.Index] = job.Value
	}

	return output
}

func main() {
	// Example usage: Transform integers by squaring them
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Apply transformation in parallel
	squared := ParallelTransform(nums, func(n int) int { return n * n })
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
