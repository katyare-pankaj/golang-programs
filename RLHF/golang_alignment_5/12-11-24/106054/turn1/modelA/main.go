package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a single task to be scheduled and executed.
type Task struct {
	ID          int
	Description string
	Duration    time.Duration
}

// Worker represents a goroutine that processes tasks from the queue.
type Worker struct {
	ID        int
	TaskQueue chan Task
	WG        *sync.WaitGroup
}

// NewWorker creates a new worker goroutine.
func NewWorker(id int, taskQueue chan Task, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		ID:        id,
		TaskQueue: taskQueue,
		WG:        wg,
	}
	worker.WG.Add(1)
	go worker.start()
	return worker
}

// start is the main function for a worker goroutine, which processes tasks from the queue.
func (w *Worker) start() {
	defer w.WG.Done()
	for task := range w.TaskQueue {
		fmt.Printf("Worker %d started task: %s\n", w.ID, task.Description)
		time.Sleep(task.Duration)
		fmt.Printf("Worker %d completed task: %s\n", w.ID, task.Description)
	}
}

func main() {
	// Create a channel to distribute tasks
	taskQueue := make(chan Task)

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 3

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		NewWorker(i, taskQueue, &wg)
	}

	// Generate and schedule tasks
	go func() {
		for i := 0; i < 10; i++ {
			task := Task{
				ID:          i,
				Description: fmt.Sprintf("Task %d", i),
				Duration:    time.Duration(rand.Intn(5)+1) * time.Second,
			}
			taskQueue <- task
		}
		close(taskQueue)
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks completed.")
}
