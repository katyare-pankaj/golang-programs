package main

import (
	"fmt"
	"sync"
)

type Task func() error

// WorkerPool struct
type WorkerPool struct {
	tasks      chan Task
	results    chan error
	numWorkers int
	wg         sync.WaitGroup
}

// NewWorkerPool initializes a new WorkerPool
func NewWorkerPool(numWorkers, taskQueueSize int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan Task, taskQueueSize),
		results:    make(chan error, taskQueueSize),
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
					wp.results <- fmt.Errorf("Worker %d: %w", workerId, err)
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
	go wp.Start()
	for _, task := range tasks {
		wp.AddTask(task)
	}
	wp.CloseTasksChannel()
	wp.CollectResults()
}

func main() {
	tasks := []Task{
		func() error { fmt.Println("Task 1"); return nil },
		func() error { fmt.Println("Task 2"); return fmt.Errorf("error in task 2") },
		func() error { fmt.Println("Task 3"); return nil },
	}

	workerPool := NewWorkerPool(3, len(tasks))
	workerPool.Run(tasks)

	for result := range workerPool.results {
		if result != nil {
			fmt.Println("Error:", result)
		}
	}
}
