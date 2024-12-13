package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Worker struct {
	ID         int
	JobChannel chan Job
	Quit       chan bool
}

func NewWorker(id int) Worker {
	return Worker{
		ID:         id,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

func (w Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job := <-w.JobChannel:
			// Simulate processing time
			fmt.Printf("Worker %d processing job %d\n", w.ID, job.ID)
			time.Sleep(time.Second) // Simulate job processing time
		case <-w.Quit:
			fmt.Printf("Worker %d stopping\n", w.ID)
			return
		}
	}
}

type WorkerPool struct {
	Workers    []Worker
	JobChannel chan Job
	Quit       chan bool
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		Workers:    make([]Worker, numWorkers),
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(i + 1)
		pool.Workers[i] = worker
		go worker.Start(&sync.WaitGroup{})
	}
	return pool
}

func (pool *WorkerPool) Start() {
	for _, worker := range pool.Workers {
		go func(w Worker) {
			for job := range pool.JobChannel {
				w.JobChannel <- job
			}
		}(worker)
	}
}

func (pool *WorkerPool) Submit(job Job) {
	pool.JobChannel <- job
}

func (pool *WorkerPool) Stop() {
	for _, worker := range pool.Workers {
		worker.Quit <- true
	}
	close(pool.JobChannel)
}

func main() {
	workerPool := NewWorkerPool(3)
	go workerPool.Start()

	for i := 1; i <= 10; i++ {
		workerPool.Submit(Job{ID: i})
	}

	time.Sleep(5 * time.Second) // Allow some time for jobs to be processed
	workerPool.Stop()
	fmt.Println("All workers stopped")
}
