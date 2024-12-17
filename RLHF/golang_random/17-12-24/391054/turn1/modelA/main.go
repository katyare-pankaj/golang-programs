package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to calculate square and simulate work
func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done on exit
	// Simulate some work with Sleep
	time.Sleep(100 * time.Millisecond)
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Timing concurrent execution
	startConcurrent := time.Now()
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)                    // Increment WaitGroup counter
		go calculateSquare(num, &wg) // Launch Goroutine
	}

	wg.Wait() // Wait for all sets of tasks to complete
	elapsedConcurrent := time.Since(startConcurrent)

	// Timing sequential execution
	startSequential := time.Now()

	for _, num := range numbers {
		calculateSquare(num, nil) // Sequential execution
	}

	elapsedSequential := time.Since(startSequential)

	fmt.Printf("Concurrent execution time: %s\n", elapsedConcurrent)
	fmt.Printf("Sequential execution time: %s\n", elapsedSequential)
}
