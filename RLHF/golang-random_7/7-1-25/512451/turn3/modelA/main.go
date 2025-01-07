package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task struct represents a task to be processed
type Task struct {
	ID    int
	Delay time.Duration
}

// ProcessTask simulates processing a task by introducing a delay
func ProcessTask(task Task, resultChan chan<- bool) {
	// Simulate task processing with a delay
	time.Sleep(task.Delay)
	// Indicate that the task is completed
	resultChan <- true
}

// Worker is a goroutine that processes tasks from the queue
func Worker(taskQueue <-chan Task, wg *sync.WaitGroup, resultChan chan<- bool) {
	defer wg.Done()
	for task := range taskQueue {
		// Start a new goroutine to process the task
		go ProcessTask(task, resultChan)
	}
}

func main() {
	// Initialize a channel for task communication
	taskQueue := make(chan Task)

	// Initialize a channel to receive task completion results
	resultChan := make(chan bool)

	// Initialize a WaitGroup to wait for all workers
	var wg sync.WaitGroup
	const numWorkers = 5 // Number of worker goroutines

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(taskQueue, &wg, resultChan)
	}

	// Generate a list of tasks
	numTasks := 20 // Number of tasks to generate
	for i := 1; i <= numTasks; i++ {
		// Generate random delay for each task
		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		taskQueue <- Task{ID: i, Delay: delay}
	}

	// Wait for all tasks to be completed
	for i := 0; i < numTasks; i++ {
		// Wait for a task completion signal
		<-resultChan
	}

	// Gracefully terminate worker goroutines
	close(taskQueue)
	wg.Wait()

	fmt.Println("All tasks completed.")
}
