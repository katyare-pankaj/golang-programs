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
	resultSize = 10000
)

type result struct {
	data int
}

// Worker represents a worker goroutine that processes tasks
type Worker struct {
	wg         *sync.WaitGroup
	tasks      <-chan int
	results    chan []int
	resultPool *sync.Pool
}

var resultPool = sync.Pool{
	New: func() interface{} {
		return &result{}
	},
}

// NewWorker creates a new worker goroutine
func NewWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- []int) *Worker {
	worker := &Worker{
		wg:      wg,
		tasks:   tasks,
		results: nil,
		resultPool: &sync.Pool{
			New: func() interface{} {
				return make([]int, resultSize)
			},
		},
	}
	return worker
}

// run starts the worker goroutine
func (w *Worker) run() {
	for task := range w.tasks {
		result := w.resultPool.Get().([]int)
		processTask(task, result)
		w.results <- result
	}
	w.wg.Done()
}

func processTask(task int, result []int) {
	// Simulate CPU-bound operation
	for i := 0; i < resultSize; i++ {
		result[i] = task * i
	}
}

func main() {
	// Set the number of GOMAXPROCS to the number of logical CPUs to optimize concurrent execution
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var wg sync.WaitGroup
	tasks := make(chan int, workItems)
	results := make(chan []int, workItems)

	// Create a pool of worker goroutines
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(&wg, tasks, results)
		wg.Add(1)
		go worker.run()
	}

	// Generate and distribute work items
	for i := 0; i < workItems; i++ {
		tasks <- i
	}
	close(tasks)

	// Collect and process results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		// Process the result here
		_ = result
		// Return the result to the pool for reuse
		resultPool.Put(result)
	}

	duration := time.Since(start)
	fmt.Printf("Completed %d tasks in %s\n", workItems, duration)
}
