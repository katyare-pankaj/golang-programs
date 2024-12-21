package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task struct represents a task to be processed
type Task struct {
	ID       int
	Name     string
	Duration time.Duration // time taken to process the task
}

// Worker simulates a worker processing a task
func Worker(task *Task, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(task.Duration)
	fmt.Printf("Worker processed task %d: %s\n", task.ID, task.Name)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Number of tasks
	numTasks := 10

	// Create a list of tasks
	tasks := make([]Task, numTasks)
	for i := 0; i < numTasks; i++ {
		tasks[i] = Task{
			ID:       i + 1,
			Name:     fmt.Sprintf("Task %d", i+1),
			Duration: time.Duration(rand.Intn(5)+1) * time.Second,
		}
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Start goroutines to process each task
	for _, task := range tasks {
		wg.Add(1) // Increment the WaitGroup count for each task
		go func(t Task) {
			defer wg.Done() // Decrement the WaitGroup count when the goroutine is complete
			Worker(&t, &wg) // Call the worker function with the task
		}(task)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("\nAll tasks have been processed successfully.")
}
