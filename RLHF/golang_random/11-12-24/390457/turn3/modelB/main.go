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

var (
	shutdownCtx, shutdownCancel = context.WithCancel(context.Background())
	wg                          = &sync.WaitGroup{}
)

func main() {
	// Set up a signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received shutdown signal, initiating graceful shutdown...")
		shutdownCancel() // Cancel the context
		wg.Wait()        // Wait for cleanup tasks to complete
		fmt.Println("Application terminated.")
		os.Exit(0)
	}()

	// Simulate an active operation (e.g., HTTP server)
	startHTTPServer(shutdownCtx)

	// Simulate a long-running goroutine
	startLongRunningTask(shutdownCtx)

	select {} // Block forever, allowing signals to handle shutdown
}

func startHTTPServer(ctx context.Context) {
	wg.Add(1)
	defer wg.Done()

	fmt.Println("HTTP server starting on :8080")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(handleRequest)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	select {
	case <-context.Background().Done():
		http.Error(w, "shutting down", http.StatusServiceUnavailable)
		return
	default:
		fmt.Fprintln(w, "Hello, World!")
	}
}

func startLongRunningTask(ctx context.Context) {
	wg.Add(1)
	defer wg.Done()

	fmt.Println("Starting long-running task...")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Long-running task shutting down...")
			return
		case <-ticker.C:
			fmt.Println("Long-running task processing...")
		}
	}
}
