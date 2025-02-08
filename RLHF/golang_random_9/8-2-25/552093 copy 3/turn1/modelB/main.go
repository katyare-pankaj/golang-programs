package main

import (
	"fmt"
	"sync"
)

// Counter is a struct that holds the count and a mutex to protect it
type Counter struct {
	count int
	mu    sync.Mutex
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	const numGoroutines = 5
	counter := Counter{}

	// Start the goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				counter.Inc()
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final count:", counter.Value())
}
