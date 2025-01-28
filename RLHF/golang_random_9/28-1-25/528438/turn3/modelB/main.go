package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Task represents a unit of work to be processed by the worker pool.
type Task struct {
	ID int
}

// worker is a function that processes tasks from the jobs channel and logs the output.
func worker(id int, jobs <-chan Task, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				// Channel closed, exit worker.
				return
			}

			start := time.Now()
			// Simulate some work with a sleep.
			time.Sleep(time.Second)
			log.Info().
				Int("worker_id", id).
				Int("task_id", job.ID).
				Dur("duration", time.Since(start)).
				Msg("Task completed")
		case <-ctx.Done():
			// Context canceled, exit worker gracefully.
			log.Info().Int("worker_id", id).Msg("Received shutdown signal, exiting worker")
			return
		}
	}
}

func main() {
	// Set up zerolog for structured logging.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create a context that can be canceled on a signal.
	ctx, cancel := context.WithCancel(context.Background())

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
		go worker(i, jobs, ctx, &wg)
	}

	// Send tasks to the workers.
	for j := 1; j <= numTasks; j++ {
		job := Task{ID: j}
		log.Info().Int("task_id", j).Msg("Queueing task")
		jobs <- job
	}

	// Close the jobs channel after sending all tasks.
	close(jobs)

	// Wait for Ctrl+C or SIGTERM signal to initiate shutdown.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until signal received.
	<-c

	// Signal the workers to exit.
	log.Info().Msg("Received shutdown signal, initiating worker exit...")
	cancel()

	// Wait for all workers to complete their graceful shutdown.
	wg.Wait()

	log.Info().Msg("All workers exited gracefully. Application shutting down.")
}
