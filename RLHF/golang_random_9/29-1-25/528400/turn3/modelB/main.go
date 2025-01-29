package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// WorkerPool represents a worker pool that manages a fixed number of goroutines.
type WorkerPool struct {
	taskChan chan int
	wg       sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with the specified number of workers.
func NewWorkerPool(numWorkers int) *WorkerPool {
	wp := &WorkerPool{
		taskChan: make(chan int),
	}
	wp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go wp.worker(i)
	}
	return wp
}

// worker is a goroutine that processes tasks from the task channel.
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.taskChan {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// AddTask adds a new task to the worker pool.
func (wp *WorkerPool) AddTask(task int) {
	wp.taskChan <- task
}

// Close shuts down the worker pool by closing the task channel and waiting for all workers to finish.
func (wp *WorkerPool) Close() {
	close(wp.taskChan)
	wp.wg.Wait()
}

func main() {
	// Create a file to store the CPU profile
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Set GOMAXPROCS to the number of available CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Number of workers in the pool
	numWorkers := 4

	// Create a new worker pool
	wp := NewWorkerPool(numWorkers)

	// Start the pprof server for memory profiling
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Add tasks to the worker pool
	totalTasks := 20
	for i := 0; i < totalTasks; i++ {
		wp.AddTask(i)
	}

	// Wait for all tasks to complete
	wp.Close()

	fmt.Println("All tasks completed.")
}
