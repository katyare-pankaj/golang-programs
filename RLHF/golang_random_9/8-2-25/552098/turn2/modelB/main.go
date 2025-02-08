package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task func()

type WorkerPool struct {
	workers  []chan Task
	tasks    chan Task
	wg       sync.WaitGroup
	shutdown chan struct{}
}

// NewWorkerPool creates a new worker pool with the specified number of workers.
func NewWorkerPool(numWorkers int) *WorkerPool {
	wp := &WorkerPool{
		workers:  make([]chan Task, numWorkers),
		tasks:    make(chan Task),
		shutdown: make(chan struct{}),
	}

	wp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		wp.workers[i] = make(chan Task)
		go wp.worker(i)
	}

	return wp
}

// Submit submits a task to the worker pool.
func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

// Stop shuts down the worker pool and waits for all workers to finish.
func (wp *WorkerPool) Stop() {
	close(wp.shutdown)
	wp.wg.Wait()
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case task := <-wp.tasks:
			task()
		case <-wp.shutdown:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

// Simulate some tasks that can be processed concurrently
func processTask1() {
	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
	fmt.Println("Task 1 completed")
}

func processTask2() {
	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
	fmt.Println("Task 2 completed")
}

func processTask3() {
	time.Sleep(time.Duration(200+rand.Intn(300)) * time.Millisecond)
	fmt.Println("Task 3 completed")
}

func main() {
	// Create a worker pool with 3 workers
	wp := NewWorkerPool(3)

	// Submit tasks to the worker pool
	wp.Submit(processTask1)
	wp.Submit(processTask2)
	wp.Submit(processTask3)
	wp.Submit(processTask1)
	wp.Submit(processTask2)
	wp.Submit(processTask3)

	// Wait for all tasks to complete
	wp.Stop()

	fmt.Println("All tasks completed")
}
