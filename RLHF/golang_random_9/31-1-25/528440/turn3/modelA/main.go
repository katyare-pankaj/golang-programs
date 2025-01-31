package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numTasks = 100000 // Number of tasks to perform
)

// work is a function that simulates a computational task.
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work with a random sleep
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func main() {
	// Configure the program to use all available CPU cores
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Start the timer
	start := time.Now()

	// Spawn all the tasks in parallel
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Stop the timer
	totalElapsed := time.Since(start)

	// Print the results
	fmt.Printf("Executed %d tasks using %d CPU cores\n", numTasks, numCPU)
	fmt.Printf("Total Execution time: %s\n", totalElapsed)

	// Analyzing Go scheduler's impact
	// The granularity of goroutine scheduling, CPU allocation, and number of CPU cores used
	// can alter how tasks are handled and the time taken to complete them.
	// In general, increasing GOMAXPROCS can improve execution time for CPU-bound tasks
	// by allowing more tasks to run concurrently.
}
