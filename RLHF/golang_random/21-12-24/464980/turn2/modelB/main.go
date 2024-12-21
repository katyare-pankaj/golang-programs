package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents the unit of work to be performed
type Task struct {
	id    int
	sleep time.Duration
}

func (t Task) run() {
	fmt.Printf("Worker %d is running task %d\n", t.id, t.id)
	time.Sleep(t.sleep)
}

// workerPool creates a worker pool with the specified number of workers.
// It takes a channel of tasks and starts a goroutine for each worker.
// Workers wait on the task channel for incoming tasks, and once a task
// is received, they process it until the channel is closed.
func workerPool(numWorkers int, taskChan <-chan Task) {
	var wg sync.WaitGroup

	// Start the specified number of workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskChan {
				task.run()
			}
		}(i + 1) // Assign worker ID starting from 1
	}

	// Wait for all workers to complete
	wg.Wait()
}

func main() {
	tasks := []Task{
		{id: 1, sleep: 2 * time.Second},
		{id: 2, sleep: 1 * time.Second},
		{id: 3, sleep: 3 * time.Second},
		{id: 4, sleep: 1 * time.Second},
		{id: 5, sleep: 2 * time.Second},
	}

	// Create a channel to send tasks to the worker pool
	taskChan := make(chan Task, len(tasks))

	// Start the worker pool with 3 workers
	go workerPool(3, taskChan)

	// Send all the tasks to the channel
	for _, task := range tasks {
		taskChan <- task
	}

	// Close the channel to signal the end of tasks
	close(taskChan)

	fmt.Println("All tasks have been completed")
}
