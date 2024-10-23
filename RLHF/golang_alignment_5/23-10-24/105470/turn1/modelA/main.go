package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numWorkers = 4
	workItems  = 1000000
)

// Worker represents a worker goroutine that processes tasks
type Worker struct {
	wg    *sync.WaitGroup
	tasks <-chan int
}

// NewWorker creates a new worker goroutine
func NewWorker(wg *sync.WaitGroup, tasks <-chan int) *Worker {
	return &Worker{wg: wg, tasks: tasks}
}

// run starts the worker goroutine
func (w *Worker) run() {
	defer w.wg.Done() // Ensure that Done is called when the worker exits
	for task := range w.tasks {
		processTask(task)
	}
}

func processTask(task int) {
	// Simulate I/O-bound operation with a sleep
	time.Sleep(time.Millisecond * 10)
	// Perform calculations or other work here
	_ = task * task
}

func main() {
	// Set the number of GOMAXPROCS to the number of logical CPUs to optimize concurrent execution
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var wg sync.WaitGroup
	tasks := make(chan int, workItems)

	// Create a pool of worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each worker
		worker := NewWorker(&wg, tasks)
		go worker.run()
	}

	// Generate and distribute work items
	for i := 0; i < workItems; i++ {
		tasks <- i
	}
	close(tasks) // Close the tasks channel to signal workers there are no more tasks

	wg.Wait() // Wait for all workers to finish

	duration := time.Since(start)
	fmt.Printf("Completed %d tasks in %s\n", workItems, duration)
}
