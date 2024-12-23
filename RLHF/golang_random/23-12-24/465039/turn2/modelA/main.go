package main

import (
	"fmt"
	"sync"
)

// Worker represents the worker function
func Worker(wg *sync.WaitGroup, start, end int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Printf("Worker %d: %d squared is %d\n", wg.Count(), i, i*i)
	}
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 3
	workerChunkSize := len(numbers) / numWorkers

	// Initialize and start workers
	for i := 0; i < numWorkers; i++ {
		startIndex := i * workerChunkSize
		endIndex := (i+1)*workerChunkSize - 1
		if i == numWorkers-1 {
			endIndex = len(numbers) - 1
		}

		wg.Add(1)
		go Worker(&wg, startIndex, endIndex)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers have completed processing.")
}
