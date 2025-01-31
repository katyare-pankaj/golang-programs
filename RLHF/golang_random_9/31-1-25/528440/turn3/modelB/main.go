package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numTasks  = 100000 // Number of tasks to be executed
	batchSize = 1000   // Number of tasks processed in a single goroutine
)

var workSizes = []int{10, 100, 1000} // Different work sizes to simulate tasks with varying CPU requirements

// Simulate computationally intensive task with a variable duration
func work(wg *sync.WaitGroup, workSize int) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(workSize)) * time.Millisecond)
}

func worker(wg *sync.WaitGroup, tasks chan int) {
	for i := range tasks {
		workSize := workSizes[i%len(workSizes)]
		work(wg, workSize)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Show the number of logical CPUs available
	numCPUs := runtime.NumCPU()
	fmt.Printf("Number of logical CPUs: %d\n", numCPUs)

	// Make the program use all available logical CPUs
	runtime.GOMAXPROCS(numCPUs)

	var wg sync.WaitGroup
	start := time.Now()

	// Create a channel for task indices
	tasks := make(chan int, numTasks)

	// Start worker goroutines
	for i := 0; i < numCPUs; i++ {
		wg.Add(1)
		go worker(&wg, tasks)
	}

	// Distribute the tasks into batches and send them to the worker goroutines
	for i := 0; i < numTasks; i += batchSize {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < batchSize; j++ {
				tasks <- i + j
			}
		}()
	}

	// Close the tasks channel to signal the end of tasks
	close(tasks)

	// Wait for all tasks and goroutines to complete
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Total Execution time: %s\n", elapsed)

	// Print the average execution time for each work size
	totalSum := map[int]time.Duration{}
	count := map[int]int{}
	for _, workSize := range workSizes {
		totalSum[workSize] = 0
		count[workSize] = 0
	}
	for i := 0; i < numTasks; i++ {
		workSize := workSizes[i%len(workSizes)]
		totalSum[workSize] += time.Duration(rand.Intn(workSize)) * time.Millisecond
		count[workSize]++
	}
	fmt.Println("Average execution time per work size:")
	for _, workSize := range workSizes {
		averageTime := totalSum[workSize] / time.Duration(count[workSize])
		fmt.Printf("Work Size: %d, Average Time: %s\n", workSize, averageTime)
	}
}
