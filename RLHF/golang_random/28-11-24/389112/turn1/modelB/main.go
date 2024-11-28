package main

import (
	"fmt"
	"sync"
	"time"
)

// workerPool represents a pool of goroutines that process work items.
type workerPool struct {
	// jobs is a channel where work items are queued.
	jobs chan func()
	// wg is used to wait for all workers to finish.
	wg sync.WaitGroup
	// numWorkers is the number of workers in the pool.
	numWorkers int
}

// newWorkerPool creates a new worker pool with the specified number of workers.
func newWorkerPool(numWorkers int) *workerPool {
	wp := &workerPool{
		jobs:       make(chan func(), 100), // Buffered channel to limit blocking
		numWorkers: numWorkers,
	}
	// Start the workers.
	wp.startWorkers()
	return wp
}

// startWorkers starts the specified number of workers.
func (wp *workerPool) startWorkers() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker runs as a goroutine and processes work items from the jobs channel.
func (wp *workerPool) worker() {
	defer wp.wg.Done()
	for job := range wp.jobs {
		job()
	}
}

// Submit submits a work item to the worker pool.
func (wp *workerPool) Submit(job func()) {
	wp.jobs <- job
}

// Stop stops the worker pool and waits for all workers to finish.
func (wp *workerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
}

// Example usage:
func main() {
	// Create a worker pool with 10 workers.
	wp := newWorkerPool(10)
	// Simulate work items being submitted to the worker pool.
	for i := 0; i < 1000; i++ {
		wp.Submit(func() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Work item completed")
		})
	}
	// Wait for all work items to be processed.
	wp.Stop()
	fmt.Println("All work items completed")
}
