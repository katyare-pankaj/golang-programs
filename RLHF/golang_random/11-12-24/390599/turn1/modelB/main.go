package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter struct with a embedded mutex to protect the count field
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment increments the counter safely
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Get returns the current count
func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// Create a new counter
	counter := &Counter{}

	// Number of goroutines to run
	numGoroutines := 1000

	// Create a channel to stop goroutines
	done := make(chan struct{})

	// Start goroutines that increment the counter
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					counter.Increment()
					time.Sleep(time.Millisecond * 100) // Simulate work
				}
			}
		}()
	}

	// Allow goroutines to run for a while
	time.Sleep(time.Second * 2)

	// Close the channel to stop goroutines
	close(done)

	// Print the final count
	fmt.Println("Final count:", counter.Get())
}
