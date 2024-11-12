package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a WorkerPool struct
type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// Start a workerPool with n workers
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
	// Example tasks
	task1 := func() {
		fmt.Println("Task 1 is running...")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 is done.")
	}

	task2 := func() {
		fmt.Println("Task 2 is running...")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 2 is done.")
	}

	// Create a worker pool with 3 workers
	wp := &WorkerPool{tasks: make(chan func())}
	wp.Start(3)

	// Add tasks to the worker pool
	wp.AddTask(task1)
	wp.AddTask(task2)
	wp.AddTask(task1)

	// Wait for all workers to finish
	wp.Wait()
}
