package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a task to be processed by the worker pool
type Task func()

// WorkerPool represents a pool of worker goroutines
type WorkerPool struct {
	workers        int            // Number of worker goroutines
	taskQueue      chan Task      // Channel for queuing tasks
	shutdownSignal chan bool      // Signal to shut down the worker pool
	wg             sync.WaitGroup // Wait group to synchronize worker shutdown
}

// NewWorkerPool creates a new worker pool with a specified number of workers
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:        workers,
		taskQueue:      make(chan Task),
		shutdownSignal: make(chan bool),
	}
}

// Start starts the worker pool
func (pool *WorkerPool) Start() {
	for i := 0; i < pool.workers; i++ {
		go pool.worker()
	}
}

// AddTask adds a task to the task queue
func (pool *WorkerPool) AddTask(task Task) {
	pool.taskQueue <- task
}

// Shutdown signals the worker pool to shut down
func (pool *WorkerPool) Shutdown() {
	pool.shutdownSignal <- true
	pool.wg.Wait()
	close(pool.taskQueue)
}

// worker is a goroutine that processes tasks from the queue
func (pool *WorkerPool) worker() {
	defer pool.wg.Done()
	for {
		select {
		case task, ok := <-pool.taskQueue:
			if !ok {
				// Queue is closed, exit worker
				return
			}
			task()
		case <-pool.shutdownSignal:
			// Shutdown signal received, exit worker
			return
		}
	}
}

func main() {
	// Create a worker pool with 5 workers
	pool := NewWorkerPool(5)
	pool.Start()

	// Generate a set of tasks to be processed
	tasks := make([]Task, 100)
	for i := range tasks {
		tasks[i] = func(id int) Task {
			return func() {
				fmt.Printf("Processing task %d...\n", id)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf("Task %d completed.\n", id)
			}
		}(i)
	}

	// Add tasks to the worker pool
	for _, task := range tasks {
		pool.AddTask(task)
	}

	// Shutdown the worker pool
	pool.Shutdown()
	fmt.Println("Worker pool shutdown complete.")
}
