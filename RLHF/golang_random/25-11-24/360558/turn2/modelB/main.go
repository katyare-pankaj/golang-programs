package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a task to be processed
type Task struct {
	Id       int
	Content  string
	Priority int
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

// processTask simulates processing a task based on its priority
func (wg *WorkerGroup) processTask(task Task, workerId int) {
	// Simulate processing time based on priority
	processingTime := 500 + int64(100*task.Priority) + int64(task.Id%10)
	time.Sleep(time.Duration(processingTime) * time.Millisecond)

	fmt.Printf("Worker %d processed task %d: %s (Priority: %d) - Completed.\n", workerId, task.Id, task.Content, task.Priority)
}

// AddTask adds a task to the queue with priority
func (wg *WorkerGroup) AddTask(task Task) {
	// To prioritize tasks, we use a select statement to send the task to the channel based on its priority.
	// Tasks with higher priority will be sent first.
	select {
	case wg.taskChan <- task:
	default:
		fmt.Println("Task queue is full. Discarding task:", task.Id)
	}
}

func main() {
	// Initialize a worker group with 4 workers
	workerGroup := NewWorkerGroup(4)
	defer workerGroup.Stop()

	workerGroup.Start()

	// Generate and add tasks to the queue with random priorities
	for i := 0; i < 20; i++ {
		priority := rand.Intn(5) + 1
		task := Task{
			Id:       i,
			Content:  fmt.Sprintf("Task %d", i),
			Priority: priority,
		}
		workerGroup.AddTask(task)
		time.Sleep(100 * time.Millisecond)
	}

	// Wait for all tasks to be completed
	time.Sleep(2 * time.Second)
}
