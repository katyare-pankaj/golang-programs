package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Initialize a slice with a range of integers
	numElements := 10000
	data := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		data[i] = i + 1
	}

	var wg sync.WaitGroup
	const numWorkers = 5

	// Start the worker goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, data, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers completed.")
	// Verify that all elements were processed
	for _, value := range data {
		if value%2 != 0 {
			fmt.Printf("Data contains unprocessed odd number: %d\n", value)
			return
		}
	}
	fmt.Println("All data elements processed successfully.")
}

func worker(workerId int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range data {
		// Simulate work by doing some operations on the data slice
		if data[i]%2 == 1 {
			data[i] *= 2
		}
		// Introduce some randomness to the work to avoid lockstep behavior
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	fmt.Printf("Worker %d completed its tasks\n", workerId)
}
