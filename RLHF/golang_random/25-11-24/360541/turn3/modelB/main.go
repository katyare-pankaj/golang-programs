package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a task as a function
type task func()

// A simple work queue with dynamic workload balance
func workQueue(tasks <-chan task, workerID int) {
	for task := range tasks {
		// Handle the task at the worker's own pace
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Worker %d completed task: %p\n", workerID, task)
	}
	fmt.Printf("Worker %d exited.\n", workerID)
}

func main() {
	// Create a buffered channel to hold tasks
	tasks := make(chan task, 100)

	// Initialize a wait group to ensure all workers finish
	var wg sync.WaitGroup

	// Define the number of worker goroutines
	numWorkers := 3

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			workQueue(tasks, workerID)
		}(i + 1)
	}

	// Generate and send tasks to the queue at a random pace
	go func() {
		for i := 0; i < 20; i++ {
			tasks <- func() {
				// Simulate work for each task
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Add random task generation delay
		}
		// Close the channel to signal workers there are no more tasks
		close(tasks)
	}()

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks are completed.")
}
