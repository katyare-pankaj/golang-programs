package main

import (
	"fmt"
	"sync"
)

// Shared resource
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment the counter safely
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Get the counter value safely
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Number of goroutines
	n := 10
	wg.Add(n)

	// Start multiple goroutines to increment the counter
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
}
