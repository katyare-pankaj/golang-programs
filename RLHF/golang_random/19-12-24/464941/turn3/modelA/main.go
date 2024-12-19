package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Shared data structure
type Counter struct {
	count int
	mu    *sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) GetCount() int {
	c.mu.RLock() // Acquire read lock
	defer c.mu.RUnlock()
	return c.count
}

func incrementCounter(c *Counter) {
	for i := 0; i < 10000; i++ {
		c.Increment()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	counter := &Counter{
		count: 0,
		mu:    &sync.Mutex{},
	}

	// Create multiple Goroutines that increment the counter
	for i := 0; i < 10; i++ {
		go incrementCounter(counter)
	}

	// Wait for all Goroutines to finish
	time.Sleep(time.Second * 2)

	fmt.Println("Final count:", counter.GetCount())
}
