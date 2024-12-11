package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID   int
	Data string
}

// WorkQueue manages the task queue
type WorkQueue struct {
	ctx         context.Context
	cancel      context.CancelFunc
	queue       chan Task
	wg          sync.WaitGroup
	workerCount int
	errorCh     chan error
}

// NewWorkQueue creates a new work queue
func NewWorkQueue(workerCount int) *WorkQueue {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkQueue{
		ctx:         ctx,
		cancel:      cancel,
		queue:       make(chan Task, 100),
		workerCount: workerCount,
		errorCh:     make(chan error),
	}
}

// Start starts the workers
func (wq *WorkQueue) Start() {
	wq.wg.Add(wq.workerCount)

	for i := 0; i < wq.workerCount; i++ {
		go wq.worker()
	}
}

// Stop stops the workers
func (wq *WorkQueue) Stop() {
	wq.cancel()
	wq.wg.Wait()
	close(wq.errorCh)
}

// Add adds a task to the queue
func (wq *WorkQueue) Add(task Task) {
	select {
	case wq.queue <- task:
	case <-wq.ctx.Done():
		log.Printf("Queue closed, unable to add task %d", task.ID)
	}
}

// worker processes tasks from the queue
func (wq *WorkQueue) worker() {
	defer wq.wg.Done()
	for {
		select {
		case task, ok := <-wq.queue:
			if !ok {
				return // Queue is closed
			}
			wq.processTask(task)
		case <-wq.ctx.Done():
			return // Context is cancelled
		}
	}
}

// processTask simulates work by sleeping and logging
func (wq *WorkQueue) processTask(task Task) {
	log.Printf("Worker %d started processing task %d with data %q", task.ID, task.ID, task.Data)
	time.Sleep(time.Duration(task.ID) * time.Second) // Simulate work
	log.Printf("Worker %d finished processing task %d", task.ID, task.ID)
}

// Errors returns a channel for receiving errors
func (wq *WorkQueue) Errors() <-chan error {
	return wq.errorCh
}

func main() {
	wq := NewWorkQueue(5)
	wq.Start()

	defer wq.Stop()

	for i := 1; i <= 10; i++ {
		wq.Add(Task{ID: i, Data: fmt.Sprintf("Task data %d", i)})
	}

	select {
	case <-time.After(5 * time.Second):
		log.Println("All tasks processed")
	case <-wq.Errors():
		log.Println("An error occurred")
	}
}
