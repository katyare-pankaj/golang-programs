package main

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Task represents a unit of work to be processed by the worker pool.
type Task struct {
	ID int
}

// worker is a function that processes tasks from the jobs channel and logs the output.
func worker(id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		start := time.Now()
		// Simulate some work with a sleep.
		time.Sleep(time.Second)
		log.Info().
			Int("worker_id", id).
			Int("task_id", job.ID).
			Dur("duration", time.Since(start)).
			Msg("Task completed")
	}
}

func main() {
	// Set up zerolog for structured logging.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create a WaitGroup for managing goroutines.
	var wg sync.WaitGroup

	// Number of workers and tasks.
	numWorkers := 3
	numTasks := 10

	// Channel for tasks.
	jobs := make(chan Task, numTasks)

	// Start the workers.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send tasks to the workers.
	for j := 1; j <= numTasks; j++ {
		job := Task{ID: j}
		log.Info().Int("task_id", j).Msg("Queueing task")
		jobs <- job
	}

	// Close the jobs channel to indicate that no more tasks will be sent.
	close(jobs)

	// Wait for all workers to complete.
	wg.Wait()

	log.Info().Msg("All tasks completed")
}
