package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work
type Task func(context.Context)

// WorkerPool manages a set of worker goroutines
type WorkerPool struct {
	ctx           context.Context
	cancel        context.CancelFunc
	workers       int
	queue         chan Task
	wg            *sync.WaitGroup
	mu            *sync.Mutex
	activeWorkers int
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		ctx:           ctx,
		cancel:        cancel,
		workers:       workers,
		queue:         make(chan Task),
		wg:            &sync.WaitGroup{},
		mu:            &sync.Mutex{},
		activeWorkers: 0,
	}
}

// Start begins processing tasks
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker is a goroutine that processes tasks from the queue
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()

	for {
		select {
		case task, ok := <-wp.queue:
			if !ok {
				return // Queue is closed, exit worker
			}

			wp.mu.Lock()
			wp.activeWorkers++
			wp.mu.Unlock()

			// Execute the task
			task(wp.ctx)

			wp.mu.Lock()
			wp.activeWorkers--
			wp.mu.Unlock()
		case <-wp.ctx.Done():
			return // Worker pool is shutting down, exit worker
		}
	}
}

// Submit adds a task to the worker pool
func (wp *WorkerPool) Submit(task Task) {
	select {
	case wp.queue <- task:
	default:
		fmt.Println("Worker pool queue is full")
	}
}

// Shutdown stops the worker pool and waits for all workers to finish
func (wp *WorkerPool) Shutdown() {
	wp.cancel()  // Signal workers to exit
	wp.wg.Wait() // Wait for all workers to finish
}

// GetActiveWorkers returns the number of active workers
func (wp *WorkerPool) GetActiveWorkers() int {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	return wp.activeWorkers
}

func main() {
	// Create a worker pool with 3 workers
	pool := NewWorkerPool(3)

	// Start the worker pool
	pool.Start()

	// Submit tasks to the worker pool
	for i := 1; i <= 10; i++ {
		fmt.Printf("Submitting task %d\n", i)
		pool.Submit(func(ctx context.Context) {
			id := i
			select {
			case <-ctx.Done():
				fmt.Printf("Task %d interrupted\n", id)
				return
			default:
				fmt.Printf("Task %d started\n", id)
				time.Sleep(time.Duration(2*id) * time.Second) // Simulate work
				fmt.Printf("Task %d completed\n", id)
			}
		})
	}

	// Wait for all tasks to complete
	time.Sleep(30 * time.Second)

	// Shut down the worker pool
	pool.Shutdown()
	fmt.Printf("All tasks completed, active workers: %d\n", pool.GetActiveWorkers())
}
