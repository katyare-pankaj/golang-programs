package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Define a task as a function
type task func()

// A worker that processes tasks from a channel
func worker(tasks chan<- task, taskCh <-chan task, finished *atomic.Bool) {
	for {
		// Get a task from the channel
		select {
		case task := <-taskCh:
			// Process the task
			go task()
		case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
			// Simulate worker doing some other work
			fmt.Printf("Worker is busy...\n")
		default:
			// If the channel is closed, we're done
			if finished.Load() {
				return
			}
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// Example task function that simulates work
func workTask() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate work
	fmt.Printf("Task completed: workTask()\n")
}

// Main function to manage workers and tasks
func main() {
	// Create a channel to hold tasks
	tasks := make(chan task)

	// Initialize a wait group to ensure all workers finish
	var wg sync.WaitGroup

	// Define the number of worker goroutines
	numWorkers := 5

	// Use an atomic boolean to signal workers to stop
	finished := atomic.NewBool(false)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, tasks, finished)
	}

	// Generate and send tasks to the queue
	for i := 0; i < 20; i++ {
		tasks <- workTask
	}

	// Wait for tasks to be processed
	close(tasks) // Close the channel to signal no more tasks
	wg.Wait()

	fmt.Println("All tasks are completed.")
}
