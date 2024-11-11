package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter represents a shared counter that will be updated by multiple goroutines
type Counter struct {
	value int64
}

// Incr increments the counter value atomically
func (c *Counter) Incr() {
	atomic.AddInt64(&c.value, 1)
}

// Decr decrements the counter value atomically
func (c *Counter) Decr() {
	atomic.AddInt64(&c.value, -1)
}

// Value returns the current counter value atomically
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	const numGoroutines = 100
	const numIterations = 10000

	var wg sync.WaitGroup
	counter := &Counter{}

	// Start multiple goroutines that will increment and decrement the counter
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIterations; j++ {
				counter.Incr()
				counter.Decr()
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Read the final counter value atomically
	finalValue := counter.Value()
	fmt.Println("Final counter value:", finalValue)

	// If we didn't use atomic operations, the final value might be non-zero
	// due to race conditions between goroutines
}
