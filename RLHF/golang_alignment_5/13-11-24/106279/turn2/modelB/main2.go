package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool represents a pool of worker goroutines
type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// Start a worker pool with n workers
func (wp *WorkerPool) Start(n int) {
	for i := 0; i < n; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for task := range wp.tasks {
				task()
			}
		}()
	}
}

// Add a task to the worker pool
func (wp *WorkerPool) AddTask(task func()) {
	wp.tasks <- task
}

// Wait for all workers to finish
func (wp *WorkerPool) Wait() {
	close(wp.tasks)
	wp.wg.Wait()
}

func main() {
	// Create a worker pool with 3 workers
	wp := &WorkerPool{tasks: make(chan func())}
	wp.Start(3)

	// Add tasks to the worker pool
	for i := 0; i < 10; i++ {
		task := func(i int) {
			fmt.Println("Processing task", i)
			time.Sleep(time.Duration(i) * time.Second)
		}(i)
		wp.AddTask(task)
	}

	// Wait for all workers to finish
	wp.Wait()
}
