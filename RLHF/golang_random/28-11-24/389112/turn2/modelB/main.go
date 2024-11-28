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

// Buffered channel to hold tasks
var taskQueue = make(chan task, 100)

// Worker function to process tasks from the queue
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		fmt.Printf("Worker %d: Processing task %d: %s\n", id, task.id, task.data)
		// Simulate work with a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	const numWorkers = 10
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}

	// Generate some sample tasks and send them to the channel
	for i := 0; i < 100; i++ {
		task := task{id: i + 1, data: fmt.Sprintf("Task %d", i+1)}
		taskQueue <- task
	}

	// Close the task channel to signal workers that no more tasks will be added
	close(taskQueue)

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed.")
}
