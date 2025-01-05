package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a struct to represent each task
type Task struct {
	id       int
	duration time.Duration // Processing time for the task
}

// Process a task with a specified duration
func processTask(task Task, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d...\n", task.id)
	// Simulate task processing time
	time.Sleep(task.duration)
}

func main() {
	var wg sync.WaitGroup
	tasks := []Task{
		{id: 1, duration: 1 * time.Second},
		{id: 2, duration: 2 * time.Second},
		{id: 3, duration: 3 * time.Second},
		{id: 4, duration: 4 * time.Second},
		{id: 5, duration: 5 * time.Second},
	}

	// Start processing tasks concurrently
	for _, task := range tasks {
		wg.Add(1)
		go processTask(task, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All tasks completed.")
}
