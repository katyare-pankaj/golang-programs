package main

import (
	"fmt"
	"log"
	"time"
)

// Process represents a job to be processed
type Process struct {
	ID int
}

// Worker represents a worker that processes jobs
type Worker struct {
	id       int
	queue    chan Process
	shutdown chan struct{}
}

// NewWorker creates a new worker
func NewWorker(id int, queue chan Process) *Worker {
	return &Worker{
		id:       id,
		queue:    queue,
		shutdown: make(chan struct{}),
	}
}

// Start starts the worker
func (w *Worker) Start() {
	defer log.Printf("Worker %d exited", w.id)

	for {
		select {
		case <-w.shutdown:
			return
		case p := <-w.queue:
			w.process(p)
		}
	}
}

// process simulates processing a job
func (w *Worker) process(p Process) {
	fmt.Printf("Worker %d processing job %d\n", w.id, p.ID)
	time.Sleep(time.Duration(p.ID%10) * 100 * time.Millisecond)
}

// shutdown stops the worker
func (w *Worker) shutdown() {
	close(w.shutdown)
}

func main() {
	numWorkers := 4
	workerQueue := make(chan Process)
	workers := make([]*Worker, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i, workerQueue)
		go workers[i].Start()
	}

	processes := []Process{
		{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10},
	}

	for _, p := range processes {
		workerQueue <- p
	}

	// Wait for all workers to finish
	for _, w := range workers {
		w.shutdown()
	}

	fmt.Println("All jobs completed.")
}
