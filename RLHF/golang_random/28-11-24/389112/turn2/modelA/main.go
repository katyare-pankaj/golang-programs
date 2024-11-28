package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the task type
type task struct {
	id   int
	data string
}

// Worker function that processes tasks from the queue
func worker(id int, tasks <-chan task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d: %s\n", id, task.id, task.data)
		// Simulate work with a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a buffered channel to hold tasks
	tasks := make(chan task, 100)

	// Start a wait group to manage goroutines
	var wg sync.WaitGroup

	// Define the number of worker goroutines
	numWorkers := 5

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, tasks, &wg)
	}

	// Generate some sample tasks and send them to the channel
	for i := 0; i < 100; i++ {
		task := task{id: i + 1, data: fmt.Sprintf("Task %d", i+1)}
		tasks <- task
	}

	// Close the task channel to signal workers that no more tasks will be added
	close(tasks)

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
