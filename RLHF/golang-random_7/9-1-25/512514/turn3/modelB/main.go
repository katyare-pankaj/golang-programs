package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the task type
type Task struct {
	// Define your task structure here
	ID   int
	Data string
}

// function to simulate processing a task
func processTask(task Task) {
	fmt.Printf("Worker %d is processing task %d\n", task.ID, task.Data)
	// Simulate processing time
	time.Sleep(time.Second * 2)
}

// worker function: receives tasks from channel and processes them
func worker(id int, tasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		processTask(task)
	}
}

func main() {
	const numWorkers = 5 // Number of worker goroutines
	var tasks chan Task

	// Create a channel to regulate the number of tasks in flight
	semaphore := make(chan struct{}, numWorkers)

	// Create a wait group to track active workers
	wg := &sync.WaitGroup{}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, wg)
	}

	// Function to generate tasks
	go func() {
		for i := 1; i <= 20; i++ {
			// Acquire semaphore to limit tasks in flight
			semaphore <- struct{}{}
			task := Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
			tasks <- task
			<-semaphore // Release semaphore once task is sent
		}
		close(tasks)
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks processed.")
}
