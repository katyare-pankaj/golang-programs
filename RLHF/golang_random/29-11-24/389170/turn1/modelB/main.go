package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Job represents a work item to be processed by a worker.
type Job struct {
    ID    int
    Data string
}

// Worker represents a worker that processes jobs from the queue.
type Worker struct {
    id       int
    queue    chan Job
    wg       *sync.WaitGroup
    isFailed bool
}

// NewWorker creates a new worker with a given id and a shared queue.
func NewWorker(id int, queue chan Job, wg *sync.WaitGroup) *Worker {
    return &Worker{
        id:    id,
        queue: queue,
        wg:    wg,
    }
}

// Start starts the worker's goroutine and begins processing jobs from the queue.
func (w *Worker) Start() {
    go w.processJobs()
}

// processJobs continuously listens for jobs on the queue and processes them.
func (w *Worker) processJobs() {
    for job := range w.queue {
        // Simulate job processing with a random delay
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

        fmt.Printf("Worker %d: Processing job %d: %s\n", w.id, job.ID, job.Data)

        // Simulate job failure with a random probability
        if rand.Intn(10) == 0 {
            fmt.Printf("Worker %d: Job %d failed\n", w.id, job.ID)
            w.isFailed = true
            break
        }
    }