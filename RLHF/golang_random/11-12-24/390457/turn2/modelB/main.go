package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	shutdownContext = context.Background()
	shutdownCancel  func()
)

func main() {
	// Create a cancellable context
	shutdownContext, shutdownCancel = context.WithCancel(shutdownContext)
	defer shutdownCancel()

	// Set up a signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received shutdown signal, initiating graceful shutdown...")
		shutdownCancel() // Cancel the context
	}()

	// Start a simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-shutdownContext.Done():
			http.Error(w, "shutting down", http.StatusServiceUnavailable)
			return
		default:
			w.Write([]byte("Hello, World!"))
		}
	})

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %v", err)
	}
}
