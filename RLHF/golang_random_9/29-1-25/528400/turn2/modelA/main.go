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

// WorkerPool manages a pool of worker goroutines
type WorkerPool struct {
	taskChan chan int
	wg       sync.WaitGroup
}

// NewWorkerPool initializes a new WorkerPool
func NewWorkerPool(numWorkers int) *WorkerPool {
	wp := &WorkerPool{taskChan: make(chan int)}
	wp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go wp.worker(i)
	}
	return wp
}

// worker processes tasks from the task channel
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.taskChan {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(100 * time.Millisecond) // Simulated work
	}
}

// AddTask sends a task to the pool
func (wp *WorkerPool) AddTask(task int) {
	wp.taskChan <- task
}

// Close shuts down the task channel and waits for all workers to finish
func (wp *WorkerPool) Close() {
	close(wp.taskChan)
	wp.wg.Wait()
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Cannot create CPU profile:", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	numWorkers := runtime.NumCPU() // Adjust to the number of logical CPUs
	wp := NewWorkerPool(numWorkers)

	for i := 0; i < 20; i++ {
		wp.AddTask(i)
	}

	// Start HTTP server to access pprof data
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wp.Close()
	fmt.Println("All tasks completed.")
}
