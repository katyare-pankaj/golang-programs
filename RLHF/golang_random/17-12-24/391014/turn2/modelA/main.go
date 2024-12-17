package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work that the worker will execute
type Task struct {
	id int
}

// Worker function that processes tasks
func worker(id int, tasks <-chan Task, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Acquire semaphore token before processing the task
		sem <- struct{}{} // Block if the channel is full
		fmt.Printf("Worker %d is processing task %d\n", id, task.id)
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Printf("Worker %d completed task %d\n", id, task.id)
		<-sem // Release semaphore token
	}
}

func main() {
	const numWorkers = 3         // Number of workers
	const maxConcurrentTasks = 2 // Max concurrent tasks (semaphore limit)

	tasks := make(chan Task, 10)                   // Channel for tasks
	sem := make(chan struct{}, maxConcurrentTasks) // Semaphore channel

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, sem, &wg)
	}

	// Send tasks to the workers
	for i := 1; i <= 10; i++ {
		tasks <- Task{id: i}
	}
	close(tasks) // Close the tasks channel to signal no more tasks

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks completed")
}
