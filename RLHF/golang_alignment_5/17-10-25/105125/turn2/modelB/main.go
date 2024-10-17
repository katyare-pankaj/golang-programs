package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work to be performed concurrently
type Task struct {
	ID   int
	Data string
}

// Worker represents a goroutine that processes tasks
type Worker struct {
	ID    int
	tasks chan Task
	quit  chan bool
	wg    sync.WaitGroup
}

// NewWorker creates a new worker goroutine
func NewWorker(id int, tasks chan Task, quit chan bool) *Worker {
	worker := &Worker{
		ID:    id,
		tasks: tasks,
		quit:  quit,
	}
	worker.wg.Add(1)
	go worker.run()
	return worker
}

// run is the main loop of the worker goroutine
func (w *Worker) run() {
	defer w.wg.Done()
	for {
		select {
		case task := <-w.tasks:
			fmt.Printf("Worker %d started processing task %d\n", w.ID, task.ID)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Worker %d completed task %d\n", w.ID, task.ID)
		case <-w.quit:
			return
		}
	}
}

// main function sets up the task queue, starts workers, and generates tasks
func main() {
	// Create a channel to send tasks to workers
	taskQueue := make(chan Task)

	// Create a channel to signal worker goroutines to quit
	quitChan := make(chan bool)

	// Start multiple worker goroutines
	const numWorkers = 5
	var workers []*Worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, NewWorker(i+1, taskQueue, quitChan))
	}

	// Generate and send tasks to the queue
	const numTasks = 20
	for i := 0; i < numTasks; i++ {
		task := Task{ID: i + 1, Data: fmt.Sprintf("Task %d", i+1)}
		taskQueue <- task
	}

	// Close the task queue to signal workers that no more tasks will be added
	close(taskQueue)

	// Wait for all workers to complete their tasks
	for _, worker := range workers {
		worker.wg.Wait()
	}

	// Close the quit channel to stop all worker goroutines
	close(quitChan)

	fmt.Println("All tasks completed.")
}
