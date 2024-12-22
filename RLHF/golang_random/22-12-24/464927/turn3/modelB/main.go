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
	// Create a WaitGroup to manage goroutines
	var wg sync.WaitGroup

	// Create a channel to receive signals
	sigCh := make(chan os.Signal, 1)

	// Listen for SIGINT and SIGTERM signals
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to perform some work
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh:
				fmt.Println("Goroutine received signal, terminating...")
				return
			default:
				fmt.Println("Goroutine is working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Start another goroutine to perform some work
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigCh:
				fmt.Println("Goroutine received signal, terminating...")
				return
			default:
				fmt.Println("Goroutine is working...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// Wait for signals to be received
	fmt.Println("Waiting for signals...")
	<-sigCh

	// Stop listening for signals
	signal.Stop(sigCh)

	// Close the signal channel
	close(sigCh)

	// Wait for goroutines to terminate gracefully
	fmt.Println("Waiting for goroutines to terminate...")
	wg.Wait()

	fmt.Println("All goroutines terminated, program exiting...")
}
