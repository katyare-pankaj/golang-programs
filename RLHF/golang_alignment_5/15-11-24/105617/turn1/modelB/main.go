package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Cache struct for LRU caching
type Cache struct {
	items map[string]int
	size  int
	max   int
	mu    sync.Mutex
}

func newCache(max int) *Cache {
	return &Cache{
		items: make(map[string]int, max),
		size:  0,
		max:   max,
	}
}

func (c *Cache) get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.items[key]
}

func (c *Cache) put(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
	if c.size >= c.max {
		for k := range c.items {
			delete(c.items, k)
			c.size--
			break
		}
	}
	c.size++
}

func main() {
	// Set the number of Goroutines to match CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create an LRU cache with a maximum size of 10000
	cache := newCache(10000)

	// Thread pool setup
	var wg sync.WaitGroup
	const numThreads = 20
	const numIterations = 100000

	// Counter for shared data access
	var sharedCounter int32

	task := func() {
		for i := 0; i < numIterations; i++ {
			key := fmt.Sprintf("key%d", i%10000)
			value := cache.get(key)
			if value == 0 {
				value = i
				cache.put(key, value)
			}
			// Simulate work
			time.Sleep(time.Nanosecond)
			// Example of concurrent shared data access
			atomic.AddInt32(&sharedCounter, 1)
		}
		wg.Done()
	}

	// Start worker threads
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go task()
	}

	wg.Wait()

	fmt.Println("Final shared counter:", sharedCounter)
}
