package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numWorkers    = 10  // Number of worker goroutines
	taskQueueSize = 100 // Size of the task queue channel
)

// Define the task type
type task struct {
	id   int
	data string
}

var (
	taskCounter   uint32 = 0
	taskProcessed uint32 = 0
)

// Worker function that processes tasks from the queue
func worker(id int, taskChan <-chan task, done chan<- bool) {
	for task := range taskChan {
		fmt.Printf("Worker %d: Processing task %d: %s\n", id, task.id, task.data)
		// Simulate work with a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		// Increment processed task counter
		atomic.AddUint32(&taskProcessed, 1)
	}
	// Signal the main loop that the worker has completed
	done <- true
}

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a buffered channel to hold tasks
	taskChan := make(chan task, taskQueueSize)
	done := make(chan bool, numWorkers)

	// Start a wait group to manage worker goroutines
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, taskChan, done)
	}

	// Generate some sample tasks and send them to the channel
	for {
		taskCounter++
		task := task{id: int(taskCounter), data: fmt.Sprintf("Task %d", int(taskCounter))}
		taskChan <- task
		// Exit the loop after sending a specified number of tasks
		if taskCounter == 100 {
			break
		}
	}

	// Close the task channel to signal workers that no more tasks will be added
	close(taskChan)

	// Wait for all worker goroutines to complete
	for i := 0; i < numWorkers; i++ {
		<-done // Blocking call until worker signals completion
	}

	// Wait for all worker goroutines to finish processing tasks
	wg.Wait()

	fmt.Printf("Total tasks processed: %d\n", taskProcessed)
	fmt.Println("All tasks completed.")
}
