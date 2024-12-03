package main

import (
	"fmt"
	"sync"
	"time"
)

// Condition Variable Example
func usingCond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	sharedResource := false

	go func() {
		mu.Lock()
		for !sharedResource {
			fmt.Println("Condition: waiting...")
			cond.Wait() // Wait for the condition to be signaled
		}
		fmt.Println("Condition: resource available!")
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second) // Simulate work

	mu.Lock()
	sharedResource = true
	cond.Signal() // Signal the waiting goroutine
	mu.Unlock()
}

// Channel Example
func usingChannel() {
	done := make(chan bool)

	go func() {
		fmt.Println("Channel: waiting for signal...")
		<-done // Wait for the signal
		fmt.Println("Channel: signal received!")
	}()

	time.Sleep(1 * time.Second) // Simulate work

	done <- true // Send signal
}

func main() {
	fmt.Println("Using Condition Variable:")
	usingCond()
	time.Sleep(2 * time.Second) // Ensure goroutine completes

	fmt.Println("\nUsing Channel:")
	usingChannel()
	time.Sleep(2 * time.Second) // Ensure goroutine completes
}
