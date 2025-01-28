package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// task represents a single unit of work to be processed by the worker pool.
type task struct {
	ID    int
	Data  string
	Delay time.Duration
}

func (t task) process() {
	time.Sleep(t.Delay)
	log.Info().Msgf("Task %d with data '%s' completed after %s", t.ID, t.Data, t.Delay)
}

func worker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan task) {
	defer wg.Done()
	for {
		select {
		case t, ok := <-tasks:
			if !ok {
				// Channel closed, worker should exit.
				log.Info().Msg("Worker exiting.")
				return
			}
			t.process()
		case <-ctx.Done():
			// Context canceled, worker should exit.
			log.Info().Msg("Worker exiting due to context cancellation.")
			return
		}
	}
}

func main() {
	// Set up structured logging using zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: stdout})

	numWorkers := runtime.NumCPU()
	fmt.Printf("Starting worker pool with %d workers\n", numWorkers)

	// Create a context that can be canceled later
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks := make(chan task, 100)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, &wg, tasks)
	}

	// Generate some example tasks and add them to the channel
	go func() {
		for i := 0; i < 100; i++ {
			t := task{
				ID:    i,
				Data:  fmt.Sprintf("Data-%d", i),
				Delay: time.Duration(i%10) * time.Second,
			}
			tasks <- t
		}
		close(tasks) // Close the channel to indicate all tasks are added
	}()

	// Wait for all workers to complete their tasks
	wg.Wait()

	log.Info().Msg("All workers done.")
}
