package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a work function that takes some amount of time to process a task
func work(taskID int, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done() // Signal that this worker has finished
	fmt.Printf("Worker #%d: Started task %d\n", taskID, taskID)
	time.Sleep(duration) // Simulate work by sleeping for the given duration
	fmt.Printf("Worker #%d: Completed task %d\n", taskID, taskID)
}

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Define a variable number of workers
	numWorkers := 3

	// Define the total number of tasks
	numTasks := 5

	// Loop over all workers and assign a task to each
	for workerID := 0; workerID < numWorkers; workerID++ {
		// Increment the WaitGroup counter for each worker
		wg.Add(1)

		// Start a goroutine for the worker
		go func() {
			defer wg.Done() // Ensure the WaitGroup counter is decremented
			for taskID := workerID; taskID < numTasks; taskID += numWorkers {
				// Each worker processes tasks in a round-robin fashion
				duration := time.Duration(taskID+1) * time.Second
				work(workerID, &wg, duration)
			}
		}()
	}

	// Wait for all workers to finish their tasks
	wg.Wait()

	fmt.Println("All tasks completed. Main function exiting.")
}
