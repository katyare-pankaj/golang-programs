package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	shutdown = make(chan struct{})
	wg       = &sync.WaitGroup{}
)

func main() {
	wg.Add(1)
	go worker()

	// Setup signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-c:
		fmt.Println("Received interrupt signal, shutting down...")
	case <-shutdown:
		fmt.Println("Worker completed, shutting down...")
	}

	// Graceful shutdown
	wg.Wait()
	close(shutdown)
	fmt.Println("Application terminated.")
}

func worker() {
	defer wg.Done()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Worker processing...")
		case <-shutdown:
			fmt.Println("Worker received shutdown signal, exiting...")
			return
		}
	}
}
