package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines = 8
	numIterations = 10000000 // Total number of iterations for the sum of squares
)

// Function to calculate the sum of squares in a given range
func sumOfSquares(start, end int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += i * i
	}
	*result += sum
}

func main() {
	// Testing different GOMAXPROCS values
	for _, procs := range []int{1, 2, 4, 8} {
		runtime.GOMAXPROCS(procs) // Set GOMAXPROCS to the desired value

		var wg sync.WaitGroup
		result := 0
		part := numIterations / numGoroutines

		startTime := time.Now() // Start timing the execution

		// Launch Goroutines to compute the sum of squares
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go sumOfSquares(i*part, (i+1)*part, &wg, &result)
		}

		wg.Wait() // Wait for all Goroutines to finish

		duration := time.Since(startTime) // Calculate duration
		fmt.Printf("GOMAXPROCS: %d, Result: %d, Duration: %s\n", procs, result, duration)
	}
}
