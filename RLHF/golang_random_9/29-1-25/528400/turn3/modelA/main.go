package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// WorkerPool manages a fixed set of worker goroutines.
// It processes tasks received on the tasks channel.
type WorkerPool struct {
	tasks     chan int
	waitGroup *sync.WaitGroup
}

// NewWorkerPool initializes the WorkerPool with a given number of workers.
func NewWorkerPool(numWorkers int, numTasks int) *WorkerPool {
	wp := &WorkerPool{
		tasks:     make(chan int, numTasks),
		waitGroup: &sync.WaitGroup{},
	}

	// Adding workers to the WaitGroup irrespective of task count as they will handle dynamic loads
	wp.waitGroup.Add(numWorkers)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go wp.worker(i)
	}

	return wp
}

// worker consumes tasks from the tasks channel and simulates work.
func (wp *WorkerPool) worker(id int) {
	defer wp.waitGroup.Done()
	for task := range wp.tasks {
		fmt.Printf("Worker %d is processing task %d\n", id, task)
		time.Sleep(time.Duration(100+rand.Int63n(900)) * time.Millisecond) // Simulate variable work time
	}
}

// AddTask submits a task to the WorkerPool.
func (wp *WorkerPool) AddTask(task int) {
	wp.tasks <- task
}

// Close shuts down the tasks channel and waits for all workers to complete.
func (wp *WorkerPool) Close() {
	close(wp.tasks)
	wp.waitGroup.Wait()
}

func main() {
	// Create a file to store the CPU profile.
	cpuProfileFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer cpuProfileFile.Close()

	// Start CPU profiling.
	pprof.StartCPUProfile(cpuProfileFile)
	defer pprof.StopCPUProfile()

	numWorkers := runtime.NumCPU()            // Set based on available CPU cores
	numTasks := 20                            // Define how many tasks you want to process.
	wp := NewWorkerPool(numWorkers, numTasks) // Create a new worker pool with the number of workers based on CPU cores.

	// Profile server using net/http/pprof for live profiling.
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Add tasks to the pool.
	for i := 0; i < numTasks; i++ {
		wp.AddTask(i)
	}

	// Complete all tasks and close the worker pool.
	wp.Close()

	fmt.Println("All tasks completed.")
}
