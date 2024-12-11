package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID int
}

// Worker function processes tasks from the queue
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Simulate work with a sleep
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		time.Sleep(1 * time.Second) // Simulate varying work duration
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Create a channel to queue tasks
	tasks := make(chan Task)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks to the queue
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}

	// Close the channel and wait for all workers to finish
	close(tasks)
	wg.Wait()

	fmt.Println("All tasks have been processed.")
}
