package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10
	taskCount  = 1000
)

var wg sync.WaitGroup

// Worker represents a worker goroutine
type Worker struct {
	workCh chan int
}

// NewWorker creates a new Worker with a given buffer size
func NewWorker(bufferSize int) *Worker {
	return &Worker{
		workCh: make(chan int, bufferSize),
	}
}

// Start starts the worker goroutine
func (w *Worker) Start() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.workCh {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate work
			fmt.Println("Worker", w.workCh, "completed task", task)
		}
	}()
}

// Stop stops the worker goroutine by closing the work channel
func (w *Worker) Stop() {
	close(w.workCh)
}

func main() {
	// Create and start worker goroutines
	workers := make([]*Worker, numWorkers)
	for i := range workers {
		workers[i] = NewWorker(100) // Buffer size
		workers[i].Start()
	}

	// Generate and send tasks
	for i := 1; i <= taskCount; i++ {
		workerIndex := rand.Intn(numWorkers)
		workers[workerIndex].workCh <- i
	}

	// Wait for all workers to complete
	wg.Wait()

	// Stop all workers
	for _, worker := range workers {
		worker.Stop()
	}
}
