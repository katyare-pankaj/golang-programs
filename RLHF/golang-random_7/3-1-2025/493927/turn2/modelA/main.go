package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a simple task to be performed
type Task struct {
	id int
}

// PerformTask simulates some work done by the task
func PerformTask(task Task, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done
	fmt.Printf("Running task %d...\n", task.id)
	time.Sleep(time.Duration(task.id) * time.Second) // Simulate different execution times
	fmt.Printf("Task %d completed.\n", task.id)
}

func main() {
	// Number of tasks
	numTasks := 5

	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Launch tasks in separate goroutines
	for i := 0; i < numTasks; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go PerformTask(Task{id: i + 1}, &wg)
	}

	// Wait for all tasks to complete before proceeding
	wg.Wait()

	fmt.Println("All tasks completed.")
}
