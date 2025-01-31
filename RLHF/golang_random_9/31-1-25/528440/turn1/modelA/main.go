package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// dummyComputeTask is a placeholder for a computational task
func dummyComputeTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Perform a dummy computation, such as calculating a large prime number.
	sum := 0.0
	for i := 0; i < 1000000; i++ {
		sum += math.Sin(float64(i))
	}
	fmt.Printf("Task %d completed: result=%f\n", id, sum)
}

func main() {
	numTasks := 10
	var wg sync.WaitGroup

	// Measure the start time
	startTime := time.Now()

	// Launch goroutines
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go dummyComputeTask(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Measure the end time
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	// Display results
	fmt.Printf("Completed %d tasks in parallel.\n", numTasks)
	fmt.Printf("Execution Time: %s\n", executionTime)
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
