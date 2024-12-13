package main

import (
	"fmt"
	"sync"
)

// Counter struct that holds the value and a mutex for synchronization
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()   // Lock the mutex before accessing the counter
	c.value++     // Increment the counter
	c.mu.Unlock() // Unlock the mutex after the operation
}

func (c *Counter) GetValue() int {
	c.mu.Lock()         // Lock the mutex for reading the value
	defer c.mu.Unlock() // Ensure the mutex is unlocked after reading
	return c.value
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the function completes
	for i := 0; i < 1000; i++ {
		c.Increment() // Increment the counter
	}
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{} // Create a new counter instance

	// Start 10 worker Goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)               // Increment the WaitGroup counter
		go worker(counter, &wg) // Start a worker Goroutine
	}

	wg.Wait() // Wait for all Goroutines to finish
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
