package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a thread-safe cache using a map and a sync.RWMutex
type Cache struct {
	data    map[int]int
	lock    sync.RWMutex
	expTime time.Duration
}

// NewCache creates a new Cache with an optional expiration time.
func NewCache(expTime ...time.Duration) *Cache {
	c := &Cache{
		data: make(map[int]int),
	}
	if len(expTime) > 0 {
		c.expTime = expTime[0]
	}
	return c
}

// Set is a thread-safe way to store data in the cache with optional expiration time.
func (c *Cache) Set(key, value int, expTime ...time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	if len(expTime) > 0 {
		expiration := time.Now().Add(expTime[0])
		go func() {
			<-time.After(expiration.Sub(time.Now()))
			c.lock.Lock()
			delete(c.data, key)
			c.lock.Unlock()
		}()
	}
}

// Get is a thread-safe way to retrieve data from the cache.
func (c *Cache) Get(key int) (value int, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	value, ok = c.data[key]
	return
}

// Delete is a thread-safe way to delete data from the cache.
func (c *Cache) Delete(key int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}

// Sample computations for demonstration
var computations = map[int]func() int{
	1: func() int { time.Sleep(1 * time.Second); return 1 * 1 },
	2: func() int { time.Sleep(2 * time.Second); return 2 * 2 },
	3: func() int { time.Sleep(3 * time.Second); return 3 * 3 },
	4: func() int { time.Sleep(4 * time.Second); return 4 * 4 },
}

func main() {
	cache := NewCache(5 * time.Second) // Cached data will expire after 5 seconds

	// Create closures to wrap computations and cache them
	for key, computation := range computations {
		func(key int, computation func() int) {
			value, ok := cache.Get(key)
			if ok {
				fmt.Printf("Got result for key %d from cache: %d\n", key, value)
			} else {
				value := computation()
				cache.Set(key, value)
				fmt.Printf("Calculated and cached result for key %d: %d\n", key, value)
			}
		}(key, computation)
	}
}
