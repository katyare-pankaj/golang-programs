package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a task to be processed.
type Task struct {
	ID int
}

// ProcessTask processes a given task by sleeping for a random duration.
func ProcessTask(task Task, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work by sleeping for a random duration
	sleepDuration := time.Duration(rand.Intn(1000)+100) * time.Millisecond
	time.Sleep(sleepDuration)

	// Print task completion
	fmt.Printf("Task %d completed after %s\n", task.ID, sleepDuration)
}

func main() {
	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a list of tasks
	tasks := make([]Task, 10)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// Start processing tasks in parallel
	for _, task := range tasks {
		wg.Add(1) // Increment the WaitGroup counter
		go ProcessTask(task, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()

	// Print a completion message
	fmt.Println("All tasks completed.")
}
