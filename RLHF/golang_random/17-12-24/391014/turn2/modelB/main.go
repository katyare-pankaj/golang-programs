package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WorkerPool struct {
	workers  []chan func()
	sem      *Semaphore
	shutdown chan struct{}
}

// NewWorkerPool creates a new worker pool with a specified number of workers
func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		workers:  make([]chan func(), size),
		sem:      NewSemaphore(size),
		shutdown: make(chan struct{}),
	}

	for i := range pool.workers {
		pool.workers[i] = make(chan func())
		go pool.workerRoutine(i)
	}

	return pool
}

func (pool *WorkerPool) workerRoutine(id int) {
	for {
		select {
		case task := <-pool.workers[id]:
			task() // Execute the task
		case <-pool.shutdown:
			return // Shut down the worker
		}
	}
}

// Submit a task to the worker pool
func (pool *WorkerPool) Submit(task func()) {
	pool.sem.Acquire()
	defer pool.sem.Release()

	id := rand.Intn(len(pool.workers))
	pool.workers[id] <- task
}

// Shutdown the worker pool
func (pool *WorkerPool) Shutdown() {
	close(pool.shutdown)
	for _, worker := range pool.workers {
		close(worker)
	}
}

func main() {
	const numWorkers = 3
	pool := NewWorkerPool(numWorkers)

	for i := 1; i <= 10; i++ {
		pool.Submit(func() {
			fmt.Printf("Task %d is starting\n", i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Task %d is done\n", i)
		})
	}

	time.Sleep(5 * time.Second)
	pool.Shutdown()
	fmt.Println("Worker pool shut down")
}
