package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// AtomicCounter struct to hold an atomic integer
type AtomicCounter struct {
	count int64
}

// Increment increments the counter atomically
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

// Get returns the current count
func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	// Create a new atomic counter
	counter := &AtomicCounter{}

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
