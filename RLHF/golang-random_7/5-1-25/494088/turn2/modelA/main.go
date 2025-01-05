package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, stop chan struct{}) {
	defer wg.Done()
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Worker stuck; exiting.")
	case <-stop:
		fmt.Println("Worker received stop signal; exiting.")
	default:
		// Simulate work
		time.Sleep(time.Duration(10 * time.Second))
	}
}

func main() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})
	wg.Add(1)

	// Start the worker goroutine
	go worker(&wg, stopCh)

	// Simulate a situation where the work takes longer than expected
	time.Sleep(3 * time.Second)

	// Decide whether to continue waiting or interrupt the worker
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Main goroutine timed out; stopping the worker.")
		close(stopCh)
		wg.Wait() // Wait for the worker to finish cleaning up
	default:
		fmt.Println("Main goroutine waiting for worker to complete.")
		wg.Wait() // Wait for the worker to finish its work
	}

	fmt.Println("Main goroutine done.")
}
