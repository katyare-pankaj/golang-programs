package main

import (
	"fmt"
	"sync"
	"time"
)

// Task is a type alias for int.  In a real application, this would be any
// type that represents the task to be processed.
type Task int

// Queue is a buffered channel that will hold tasks to be processed.
type Queue chan Task

// Worker represents a single worker process that will take tasks from the queue and process them.
type Worker struct {
	id    int
	queue Queue
	wg    *sync.WaitGroup
}

// NewWorker creates a new Worker instance.
func NewWorker(id int, queue Queue, wg *sync.WaitGroup) Worker {
	return Worker{
		id:    id,
		queue: queue,
		wg:    wg,
	}
}

// Run starts the worker's goroutine, which will process tasks from the queue until the queue is closed.
func (w Worker) Run() {
	defer w.wg.Done()
	for task := range w.queue {
		fmt.Printf("Worker %d is processing task %d\n", w.id, task)
		processTask(task)
	}
	fmt.Printf("Worker %d has finished\n", w.id)
}

// processTask simulates processing a task by sleeping for a random amount of time.
func processTask(task Task) {
	duration := time.Duration(time.Millisecond * time.Duration(100*(task%5)))
	time.Sleep(duration)
}

// main creates a task queue with 10 worker processes and feeds them a series of tasks.
// It then waits for all of the workers to finish before exiting the program.
func main() {
	queue := make(Queue, 100) // Buffered channel with capacity 100
	var wg sync.WaitGroup

	// Start 10 worker processes
	for i := 0; i < 10; i++ {
		worker := NewWorker(i+1, queue, &wg)
		wg.Add(1)
		go worker.Run()
	}

	// Feed tasks to the queue
	for i := 1; i <= 20; i++ {
		queue <- Task(i)
	}

	close(queue) // Close the queue to signal that no more tasks will be added
	wg.Wait()    // Wait for all workers to finish
	fmt.Println("All tasks have been processed")
}
