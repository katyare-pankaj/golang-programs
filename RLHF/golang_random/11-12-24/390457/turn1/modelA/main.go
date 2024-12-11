package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Channel to listen for OS signals
	sigs := make(chan os.Signal, 1)
	// Channel to notify when we're ready to exit
	done := make(chan bool, 1)

	// Notify about SIGINT and SIGTERM signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Use WaitGroup for concurrent operations
	var wg sync.WaitGroup

	// Start a goroutine to simulate some work
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("Goroutine terminating gracefully...")
				return
			default:
				fmt.Println("Doing some work...")
				time.Sleep(1 * time.Second) // Simulate some work with sleep
			}
		}
	}()

	// Wait for a signal
	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %s\n", sig)
		done <- true // Notify the goroutine to finish
	}()

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("Exiting application gracefully...")
}