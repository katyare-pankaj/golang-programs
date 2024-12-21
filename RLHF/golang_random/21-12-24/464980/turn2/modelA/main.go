package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work to be performed by a worker.
type Task struct {
	id int
}

// WorkerPool is a structure that manages a pool of worker Goroutines.
type WorkerPool struct {
	numWorkers int
	tasks      chan Task
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers.
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan Task),
		wg:         sync.WaitGroup{},
	}
}

// Start starts the worker pool, launching the specified number of worker Goroutines.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker is a Goroutine that processes tasks from the channel.
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.tasks {
		fmt.Printf("Worker %d processing task %d\n", wp.numWorkers, task.id)
		// Simulate work
		time.Sleep(time.Second)
	}
}

// Stop stops the worker pool by closing the task channel and waiting for all workers to finish.
func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
	fmt.Println("Worker pool stopped")
}

// Submit submits a task to the worker pool for processing.
func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

func main() {
	// Create a worker pool with 3 workers
	wp := NewWorkerPool(3)
	wp.Start()

	// Submit some tasks
	for i := 0; i < 10; i++ {
		wp.Submit(Task{id: i + 1})
	}

	// Wait for the worker pool to stop
	wp.Stop()
}
