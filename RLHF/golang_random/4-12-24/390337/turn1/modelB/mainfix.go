package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Worker struct {
	id       int
	queue    chan Job
	wg       *sync.WaitGroup
	isFailed bool
}

func NewWorker(id int, queue chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:    id,
		queue: queue,
		wg:    wg,
	}
}

func (w *Worker) Start() {
	go w.processJobs()
}

func (w *Worker) processJobs() {
	defer w.wg.Done() // Decrement the counter when the goroutine finishes

	for job := range w.queue {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		fmt.Printf("Worker %d: Processing job %d: %s\n", w.id, job.ID, job.Data)

		if rand.Intn(10) == 0 {
			fmt.Printf("Worker %d: Job %d failed\n", w.id, job.ID)
			w.isFailed = true
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	jobQueue := make(chan Job, 10) // Create a buffered channel for job queue
	var wg sync.WaitGroup

	// Create workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, jobQueue, &wg)
		worker.Start()
	}

	// Dispatch jobs
	numJobs := 5
	for i := 1; i <= numJobs; i++ {
		jobQueue <- Job{ID: i, Data: fmt.Sprintf("Job data %d", i)}
	}

	close(jobQueue) // Close the queue so workers know to stop

	// Wait for all workers to finish
	wg.Add(numWorkers)
	wg.Wait()

	fmt.Println("All jobs processed.")
}
