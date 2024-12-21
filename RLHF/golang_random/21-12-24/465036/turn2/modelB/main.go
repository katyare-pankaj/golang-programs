package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// Task represents a work item to be processed
type Task struct {
	ID       int
	Duration time.Duration // Time taken to complete the task
}

// Worker is a function that takes a Task and processes it
type Worker func(task Task)

func main() {
	// Number of workers to create
	const numWorkers = 3

	// Create a buffered channel to hold tasks
	tasks := make(chan Task, 100)

	// Create a WaitGroup to track the number of tasks in flight
	var wg sync.WaitGroup

	// Launch workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, &wg)
	}

	// Generate and submit tasks
	for i := 1; i <= 10; i++ {
		task := Task{
			ID:       i,
			Duration: time.Duration(rand.Intn(5)+1) * time.Second,
		}
		wg.Add(1)
		tasks <- task // Send the task to the worker channel
	}

	// Close the channel to signal workers that there are no more tasks
	close(tasks)

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("All tasks completed.")
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d: Starting task %d\n", id, task.ID)
		time.Sleep(task.Duration) // Simulate task processing
		fmt.Printf("Worker %d: Completed task %d\n", id, task.ID)
		wg.Done()
	}
}
