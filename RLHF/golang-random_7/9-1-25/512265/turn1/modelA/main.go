package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch <-chan int, result chan<- int) {
	for num := range ch {
		// Simple processing function
		result <- num * num
	}
	close(result)
}

func processSliceConcurrently(numbers []int, numWorkers int) []int {
	const chunkSize = 100 // Adjust chunk size based on workload and hardware

	var wg sync.WaitGroup
	channels := make([]chan int, numWorkers)
	results := make([]chan int, numWorkers)

	// Prepare worker channels
	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan int, chunkSize*10)
		results[i] = make(chan int)
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			worker(channels[idx], results[idx])
		}(i)
	}

	// Divide work into chunks and send to workers
	for i, num := range numbers {
		workerIndex := i % numWorkers
		channels[workerIndex] <- num
	}

	// Close channels to signal workers
	for i := 0; i < numWorkers; i++ {
		close(channels[i])
	}

	var processedNumbers []int
	for _, result := range results {
		for range result {
			processedNumbers = append(processedNumbers, <-result)
		}
	}

	wg.Wait()
	return processedNumbers
}

func main() {
	start := time.Now()

	// Test slice with a large number of integers
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}

	// Number of worker goroutines
	numWorkers := 4

	// Process numbers concurrently
	processed := processSliceConcurrently(numbers, numWorkers)

	// Check and print the result
	if len(processed) == len(numbers) {
		fmt.Println("Processing completed successfully!")
	} else {
		fmt.Println("Error: Processing did not return the expected number of results.")
	}

	fmt.Printf("Processing time: %v\n", time.Since(start))
}
