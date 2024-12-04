package main

import (
	"fmt"
	"sync"
)

// Task type represents a function that returns an error
type Task func() error

// WorkerPool struct
type WorkerPool struct {
	tasks      chan Task
	results    chan string
	numWorkers int
	wg         sync.WaitGroup
}

// NewWorkerPool initializes a new WorkerPool
func NewWorkerPool(numWorkers, taskQueueSize int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan Task, taskQueueSize),
		results:    make(chan string, taskQueueSize),
		numWorkers: numWorkers,
	}
}

// Start initializes worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func(workerId int) {
			defer wp.wg.Done()
			for task := range wp.tasks {
				err := task()
				if err != nil {
					wp.results <- fmt.Sprintf("Worker %d: Task failed with error: %s", workerId, err)
				} else {
					wp.results <- fmt.Sprintf("Worker %d: Task succeeded", workerId)
				}
			}
		}(i)
	}
}

// AddTask adds a task to the pool
func (wp *WorkerPool) AddTask(task Task) {
	wp.tasks <- task
}

// CloseTasksChannel closes the task channel
func (wp *WorkerPool) CloseTasksChannel() {
	close(wp.tasks)
}

// CollectResults collects task results and waits for all workers to finish
func (wp *WorkerPool) CollectResults() {
	wp.wg.Wait()
	close(wp.results)
}

// Run demonstrates how to use the WorkerPool
func (wp *WorkerPool) Run(tasks []Task) {
	wp.Start()
	for _, task := range tasks {
		wp.AddTask(task)
	}
	wp.CloseTasksChannel()
	wp.CollectResults()
}

func main() {
	// Define tasks
	tasks := []Task{
		func() error { fmt.Println("Processing Task 1"); return nil },
		func() error { fmt.Println("Processing Task 2"); return fmt.Errorf("error in task 2") },
		func() error { fmt.Println("Processing Task 3"); return nil },
	}

	// Create worker pool
	workerPool := NewWorkerPool(3, len(tasks))

	// Run tasks
	go workerPool.Run(tasks)

	// Print results
	for result := range workerPool.results {
		fmt.Println(result)
	}
}
