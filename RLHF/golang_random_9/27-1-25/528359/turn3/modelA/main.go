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
func worker(ctx context.Context, id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Info().Int("worker_id", id).Msg("Worker shutting down")
			return
		case job, ok := <-jobs:
			if !ok {
				log.Info().Int("worker_id", id).Msg("No more jobs, worker exiting")
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
		}
	}
}

func main() {
	// Set up zerolog for structured logging.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create a channel to listen for interrupt or terminate signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Create a context that will be canceled when a shutdown signal is received.
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
		go worker(ctx, i, jobs, &wg)
	}

	// Send tasks to the workers.
	for j := 1; j <= numTasks; j++ {
		job := Task{ID: j}
		log.Info().Int("task_id", j).Msg("Queueing task")
		jobs <- job
	}

	// Wait for an interruption signal.
	<-stop
	log.Info().Msg("Shutdown signal received")

	// Cancel the context to signal goroutines to stop.
	cancel()

	// Close the jobs channel to prevent any more tasks from being sent.
	close(jobs)

	// Wait for all workers to complete.
	wg.Wait()

	log.Info().Msg("All tasks completed and all workers shut down gracefully")
}
