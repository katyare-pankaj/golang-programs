package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be processed
type Task struct {
	Id       int
	Content  string
	Deadline time.Time
}

// WorkerGroup manages a set of worker goroutines
type WorkerGroup struct {
	workers      int
	taskChan     chan Task
	wg           sync.WaitGroup
	shutdownChan chan struct{}
}

// NewWorkerGroup initializes a new WorkerGroup
func NewWorkerGroup(workers int) *WorkerGroup {
	return &WorkerGroup{
		workers:      workers,
		taskChan:     make(chan Task),
		shutdownChan: make(chan struct{}),
	}
}

// Start starts the worker goroutines
func (wg *WorkerGroup) Start() {
	for i := 0; i < wg.workers; i++ {
		go wg.runWorker(i)
	}
}

// Stop gracefully stops the worker goroutines
func (wg *WorkerGroup) Stop() {
	close(wg.shutdownChan)
	wg.wg.Wait()
}

// runWorker runs a single worker goroutine
func (wg *WorkerGroup) runWorker(id int) {
	wg.wg.Add(1)
	defer wg.wg.Done()

	for {
		select {
		case task, ok := <-wg.taskChan:
			if !ok {
				return // Channel closed, worker exits
			}
			// Process the task
			wg.processTask(task, id)

		case <-wg.shutdownChan:
			return // Shutdown signal received, worker exits
		}
	}
}

// processTask simulates processing a task
func (wg *WorkerGroup) processTask(task Task, workerId int) {
	fmt.Printf("Worker %d processing task %d: %s\n", workerId, task.Id, task.Content)

	// Simulate processing time
	time.Sleep(time.Duration(500+int64(task.Id%10)) * time.Millisecond)

	// Check if task is timely
	if time.Now().After(task.Deadline) {
		fmt.Printf("Task %d processed by Worker %d - TIMEOUT!\n", task.Id, workerId)
	} else {
		fmt.Printf("Task %d processed by Worker %d - Completed.\n", task.Id, workerId)
	}
}

// AddTask adds a task to the queue
func (wg *WorkerGroup) AddTask(task Task) {
	wg.taskChan <- task
}

func main() {
	workerGroup := NewWorkerGroup(4)
	defer workerGroup.Stop()

	workerGroup.Start()

	// Generate and add tasks to the queue
	for i := 0; i < 20; i++ {
		time.Sleep(50 * time.Millisecond)
		deadline := time.Now().Add(500 * time.Millisecond)
		workerGroup.AddTask(Task{Id: i, Content: fmt.Sprintf("Task %d", i), Deadline: deadline})
	}

	// Wait for the queue to empty
	time.Sleep(2 * time.Second)
}
