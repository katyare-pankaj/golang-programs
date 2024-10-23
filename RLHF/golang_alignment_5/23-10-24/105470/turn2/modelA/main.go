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

// Worker represents a worker goroutine
type Worker struct {
	wg     *sync.WaitGroup
	tasks  <-chan *task
	result chan<- *result
}

type task struct {
	data int
}

type result struct {
	data int
}

var taskPool = sync.Pool{
	New: func() interface{} {
		return &task{}
	},
}

var resultPool = sync.Pool{
	New: func() interface{} {
		return &result{}
	},
}

// NewWorker creates a new worker goroutine
func NewWorker(wg *sync.WaitGroup, tasks <-chan *task, result chan<- *result) *Worker {
	worker := &Worker{wg: wg, tasks: tasks, result: result}
	return worker
}

// run starts the worker goroutine
func (w *Worker) run() {
	defer w.wg.Done() // Ensure that Done is called when the goroutine exits
	for t := range w.tasks {
		r := processTask(t)
		w.result <- r
		// Recycle task and result objects
		taskPool.Put(t)
		resultPool.Put(r)
	}
}

func processTask(t *task) *result {
	// Simulate work
	time.Sleep(time.Millisecond * 1)
	r := resultPool.Get().(*result)
	r.data = t.data * t.data
	return r
}

func main() {
	// Set GOMAXPROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	var wg sync.WaitGroup
	tasks := make(chan *task, workItems)
	results := make(chan *result, workItems)

	// Create workers
	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(&wg, tasks, results)
		wg.Add(1)
		go worker.run()
	}

	// Generate tasks
	for i := 0; i < workItems; i++ {
		t := taskPool.Get().(*task)
		t.data = i
		tasks <- t
	}
	close(tasks)

	// Collect results
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(results)
	}()

	count := 0
	for r := range results {
		count += r.data
		// Recycle result object
		resultPool.Put(r)
	}

	duration := time.Since(start)
	fmt.Printf("Completed %d tasks, sum: %d in %s\n", workItems, count, duration)
}
