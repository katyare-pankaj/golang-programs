package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a task struct to hold task identifiers and results
type task struct {
	id     int
	result string
}

// RunTask simulates performing some work and returns a result
func RunTask(taskID int, duration time.Duration) string {
	time.Sleep(duration) // Simulate work taking time
	return fmt.Sprintf("Task %d finished with result: result%d", taskID, taskID)
}

// Worker function picks tasks from the channel, processes them, and sends results back
func Worker(taskCh chan<- task, resultCh chan<- task) {
	for {
		// Get a task from the channel
		taskToProcess, ok := <-taskCh
		if !ok {
			// Channel closed, break out of loop
			break
		}
		// Simulate work and generate result
		result := RunTask(taskToProcess.id, time.Duration(taskToProcess.id)*time.Second)
		// Send result back through result channel
		resultCh <- task{taskToProcess.id, result}
	}
}

// Main function sets up the worker pool, processes tasks, and summarizes results
func main() {
	var wg sync.WaitGroup
	// Number of worker threads
	numWorkers := 5

	// Channels for task distribution and result collection
	taskCh := make(chan<- task, numWorkers)
	resultCh := make(chan task, numWorkers)

	// Start worker threads
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			Worker(taskCh, resultCh)
			wg.Done()
		}()
	}

	// Add independent tasks
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, taskID := range tasks {
		taskCh <- task{taskID, ""}
	}

	// Close the task channel to signal all tasks have been distributed
	close(taskCh)

	// Collect results from all workers
	results := make([]task, len(tasks))
	for i, _ := range tasks {
		wg.Add(1)
		go func(i int) {
			results[i] = <-resultCh
			wg.Done()
		}(i)
	}

	// Wait for all workers to finish and collect results
	wg.Wait()

	// Summarize results
	fmt.Println("Task results:")
	for _, result := range results {
		fmt.Println(result)
	}
}
