package main

import (
	"fmt"
	"runtime"
	"sync"
)

func parallelSum(data []int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	sum := 0
	for _, num := range data {
		sum += num
	}
	results <- sum
}

func main() {
	fmt.Println("Go Concurrency Sum Calculation")

	// Simulate a large integer slice
	sliceSize := 1000000
	data := make([]int, sliceSize)
	for i := 0; i < sliceSize; i++ {
		data[i] = i + 1
	}

	// Calculate the number of available cores for concurrency
	numCores := runtime.NumCPU()
	fmt.Printf("Number of cores: %d\n", numCores)

	// Determine the size of each segment
	segmentSize := sliceSize / numCores
	fmt.Printf("Size of each segment: %d\n", segmentSize)

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Create a channel to receive the results from each goroutine
	results := make(chan int, numCores)

	// Start goroutines to calculate the sum of each segment concurrently
	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		start := i * segmentSize
		end := (i + 1) * segmentSize
		if end > sliceSize {
			end = sliceSize
		}
		go parallelSum(data[start:end], &wg, results)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)

	// Calculate the final sum by aggregating the results from each goroutine
	totalSum := 0
	for result := range results {
		totalSum += result
	}

	fmt.Printf("Total sum of all elements: %d\n", totalSum)
}
