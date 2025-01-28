package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work with a unique ID and a simulated processing time.
type Task struct {
	ID       int
	Duration time.Duration
}

// worker simulates a worker that processes tasks from the taskCh channel.
func worker(id int, taskCh <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that the wait group is decremented when the worker finishes.
	for task := range taskCh {
		fmt.Printf("Worker %d: Starting task %d\n", id, task.ID)
		time.Sleep(task.Duration) // Simulate the task processing.
		fmt.Printf("Worker %d: Completed task %d\n", id, task.ID)
	}
}

func main() {
	// Define constants for the number of workers and tasks.
	const numWorkers = 5
	const numTasks = 20

	// Create a channel to send tasks to the workers.
	taskCh := make(chan Task)

	// Initialize a wait group to synchronize the completion of tasks.
	var wg sync.WaitGroup

	// Start worker goroutines.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Add a count for each worker to the wait group.
		go worker(i+1, taskCh, &wg)
	}

	// Generate tasks and send them to the taskCh channel.
	go func() {
		for i := 0; i < numTasks; i++ {
			task := Task{
				ID:       i + 1,
				Duration: time.Duration(rand.Intn(2000)+500) * time.Millisecond, // Random duration between 500ms and 2500ms
			}
			taskCh <- task
		}
		close(taskCh) // Close the channel after sending all tasks.
	}()

	// Wait for all workers to complete their tasks.
	wg.Wait()
	fmt.Println("All tasks have been completed.")
}
