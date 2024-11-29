package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    minWorkers        = 2           // Minimum number of workers in the pool
    maxWorkers        = 10          // Maximum number of workers in the pool
    maxRetryAttempts  = 3          // Maximum number of retry attempts for failed jobs
    workerIdleTimeout = 1 * time.Second // Timeout for worker idle
)

// Job represents a work item to be processed by a worker.
type Job struct {
    ID    int
    Data string
    RetryCount int  // To track retry attempts
}

// Worker represents a worker that processes jobs from the queue.
type Worker struct {
    id           int
    jobQueue     chan Job
    retryQueue   chan Job
    shutdownChan chan struct{}
    wg           *sync.WaitGroup
}

// NewWorker creates a new worker with the given id and shared job and retry queues.
func NewWorker(id int, jobQueue, retryQueue chan Job, shutdownChan chan struct{}, wg *sync.WaitGroup) *Worker {
    return &Worker{
        id:           id,
        jobQueue:     jobQueue,
        retryQueue:   retryQueue,
        shutdownChan: shutdownChan,
        wg:           wg,
    }
}

// Start starts the worker's goroutine and begins processing jobs from the queue.
func (w *Worker) Start() {
    go w.processJobs()
}

// processJobs continuously listens for jobs on the job queue and processes them.
// If a job fails, it is retried after an exponential backoff period.
func (w *Worker) processJobs() {
    defer w.wg.Done()

    for {
        select {
        case job := <-w.jobQueue:
            fmt.Printf("Worker %d: Processing job %d: %s\n", w.id, job.ID, job.Data)

            // Simulate job processing with a random delay
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

            // Simulate job failure with a random probability
            if rand.Intn(10) == 0 {
                fmt.Printf("Worker %d: Job %d failed\n", w.id, job.ID)
                // Retry the job if it hasn't exceeded the maximum retry attempts
                if job.RetryCount < maxRetryAttempts {
                    job.RetryCount++
                    w.retryQueue <- job
                } else {
                    fmt.Printf("Worker %d: Job %d failed after %d retries. Giving up.\n", w.id, job.ID, job.RetryCount)
                }
                break
            }

            // If the job succeeds, no need to retry it

        case <-w.shutdownChan:
            fmt.Printf("Worker %d: Received shutdown signal. Shutting down.\n", w.id)
            return
        }
    }
}

// WorkerPool represents the dynamic worker pool manager.
type WorkerPool struct {
    jobQueue          chan Job
    retryQueue        chan Job
    workers           []*Worker
    workerShutdownChans []chan struct{}
    numWorkers        int
    wg                *sync.WaitGroup