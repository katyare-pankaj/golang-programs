package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a task as a function
type task func()

// A simple task queue implementation using a channel
func taskQueue(tasks chan<- task, numWorkers int, wg *sync.WaitGroup) {
	for {
		// Get a task from the channel
		select {
		case task := <-tasks:
			// Process the task
			go task()
		default:
			// If the channel is closed, we're done
			wg.Done()
			return
		}
	}
}

// Example task function that simulates work
func workTask() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate work
	fmt.Printf("Task completed: workTask()\n")
}

func main() {
	// Create a channel to hold tasks
	tasks := make(chan task)

	// Initialize a wait group to ensure all workers finish
	var wg sync.WaitGroup

	// Define the number of worker goroutines
	numWorkers := 5

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go taskQueue(tasks, numWorkers, &wg)
	}

	// Generate and send tasks to the queue
	for i := 0; i < 20; i++ {
		tasks <- workTask
	}

	// Close the channel when all tasks are sent
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All tasks are completed.")
}
