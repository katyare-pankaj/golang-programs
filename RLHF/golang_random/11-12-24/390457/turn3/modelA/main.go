package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// A simple data structure for simulating background work
type worker struct {
	id int
}

// Start a worker that simulates some long-running task
func (w *worker) start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when it returns

	for {
		select {
		case <-ctx.Done():
			// Cleanup before exiting the worker
			fmt.Printf("Worker %d stopping gracefully...\n", w.id)
			return
		default:
			// Simulate work
			fmt.Printf("Worker %d is working...\n", w.id)
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}
}

func main() {
	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Create a cancellable context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Set up signal handling for SIGINT and SIGTERM
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	// Start workers
	workerCount := 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		w := &worker{id: i + 1}
		go w.start(ctx, &wg)
	}

	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Start the HTTP server
	go func() {
		fmt.Println("Starting server at :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-sigChannel
	fmt.Println("Received shutdown signal, initiating graceful shutdown...")
	cancel() // Cancel the context

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers have completed. Shutting down...")
}
