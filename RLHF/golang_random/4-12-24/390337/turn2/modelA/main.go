package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a work item to be processed by a worker.
type Job struct {
	ID   int
	Data string
}

// WorkerPool manages a pool of workers.
type WorkerPool struct {
	jobQueue       chan Job
	maxWorkers     int
	currentWorkers int
	idleTimeout    time.Duration
	mu             sync.Mutex
	wg             *sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with specified max workers and idle timeout.
func NewWorkerPool(maxWorkers int, idleTimeout time.Duration) *WorkerPool {
	return &WorkerPool{
		jobQueue:       make(chan Job),
		maxWorkers:     maxWorkers,
		idleTimeout:    idleTimeout,
		currentWorkers: 0,
		wg:             &sync.WaitGroup{},
	}
}

// AddJob submits a job to the worker pool.
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobQueue <- job
}

// Start starts the worker pool, dynamically managing the number of workers.
func (wp *WorkerPool) Start() {
	go wp.manageWorkers()
}

// manageWorkers dynamically adjusts the number of workers based on the workload.
func (wp *WorkerPool) manageWorkers() {
	for {
		wp.mu.Lock()
		workload := len(wp.jobQueue)
		idleWorkers := wp.currentWorkers - workload

		if workload > 0 {
			if wp.currentWorkers < wp.maxWorkers {
				wp.currentWorkers++
				wp.mu.Unlock()
				wp.wg.Add(1)
				go wp.worker()
			}
		} else if idleWorkers > 0 {
			wp.currentWorkers--
			wp.mu.Unlock()
			wp.wg.Done()
		} else {
			wp.mu.Unlock()
		}

		if idleWorkers == wp.currentWorkers {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// worker processes jobs from the queue and handles retries.
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	var lastJobProcessed time.Time

	for {
		select {
		case job := <-wp.jobQueue:
			lastJobProcessed = time.Now()
			err := wp.processJob(job)
			if err != nil {
				fmt.Printf("Job %d failed, retrying...\n", job.ID)
				wp.jobQueue <- job
			}
		default:
			if time.Since(lastJobProcessed) > wp.idleTimeout {
				fmt.Println("Worker idle, shutting down.")
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// processJob simulates processing a job with a random delay and failure probability.
func (wp *WorkerPool) processJob(job Job) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	fmt.Printf("Processing job %d: %s\n", job.ID, job.Data)

	// Simulate job failure with a random probability
	if rand.Intn(10) == 0 {
		return fmt.Errorf("job %d failed", job.ID)
	}

	return nil
}

// Wait waits for all workers to finish processing jobs.
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wp := NewWorkerPool(10, 5*time.Second)
	wp.Start()

	// Generate and add jobs
	for i := 0; i < 50; i++ {
		wp.AddJob(Job{ID: i, Data: fmt.Sprintf("Job %d", i)})
		time.Sleep(100 * time.Millisecond)
	}

	// Wait for all workers to finish
	wp.Wait()

	fmt.Println("Worker pool shut down.")
}
