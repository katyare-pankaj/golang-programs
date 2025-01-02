package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the type of task that will be processed by the worker pool
type Task struct {
	ID   int
	Data string
}

// Worker function that processes tasks from the channel
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started processing task %d: %s\n", id, task.ID, task.Data)
		// Simulate work by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Worker %d finished processing task %d\n", id, task.ID)
	}
}

func main() {
	// Number of workers in the pool
	numWorkers := 4

	// Create a buffered channel to hold tasks
	tasks := make(chan Task, 100)

	// Create a wait group to synchronize the main goroutine with the workers
	var wg sync.WaitGroup

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Generate a large number of tasks and send them to the channel
	for i := 1; i <= 100; i++ {
		task := Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
		tasks <- task
	}

	// Close the channel to signal the workers that there are no more tasks
	close(tasks)

	// Wait for all workers to finish processing tasks
	wg.Wait()

	fmt.Println("All tasks completed.")
}
