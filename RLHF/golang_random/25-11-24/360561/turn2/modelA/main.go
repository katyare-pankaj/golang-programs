package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// Task represents a task to be processed
type Task struct {
	Id       int
	Content  string
	Priority int // Higher priority means higher urgency
}

// WorkerGroup manages a set of worker goroutines
type WorkerGroup struct {
	workers      int
	taskQueue    []Task
	taskChan     chan Task
	wg           sync.WaitGroup
	shutdownChan chan struct{}
	mu           sync.Mutex
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
			wg.processTask(task, id)

		case <-wg.shutdownChan:
			return // Shutdown signal received, worker exits
		}
	}
}

// processTask simulates processing a task
func (wg *WorkerGroup) processTask(task Task, workerId int) {
	fmt.Printf("Worker %d processing task %d with priority %d: %s\n", workerId, task.Id, task.Priority, task.Content)

	// Simulate processing time
	time.Sleep(time.Duration(500+int64(task.Id%10)) * time.Millisecond)
}

// AddTask adds a task to the queue
func (wg *WorkerGroup) AddTask(task Task) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.taskQueue = append(wg.taskQueue, task)
	wg.taskQueue = sortTasksByPriority(wg.taskQueue)

	wg.taskChan <- task
}

// sortTasksByPriority sorts tasks in descending order of priority
func sortTasksByPriority(tasks []Task) []Task {
	return sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})
}

func main() {
	workerGroup := NewWorkerGroup(4)
	defer workerGroup.Stop()

	workerGroup.Start()

	// Generate and add tasks to the queue
	for i := 0; i < 20; i++ {
		priority := i%3 + 1 // Priorities range from 1 to 3
		workerGroup.AddTask(Task{Id: i, Content: fmt.Sprintf("Task %d", i), Priority: priority})
		time.Sleep(50 * time.Millisecond)
	}

	// Wait for the queue to empty
	time.Sleep(2 * time.Second)
}
