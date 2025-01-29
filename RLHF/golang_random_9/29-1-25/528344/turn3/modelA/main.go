package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a struct that holds a map for cached values and a mutex for thread safety.
type Cache struct {
	data map[string]interface{}
	mu   sync.Mutex
}

// NewCache creates and returns a new Cache instance.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Get retrieves a value from the cache for a given key.
// It returns the value and a boolean indicating whether the key was found.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.data[key]
	return value, found
}

// Set adds a key-value pair to the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// ThreadSafeCacheExample demonstrates using the Cache in a concurrent context.
func ThreadSafeCacheExample() {
	cache := NewCache()

	// Simulate concurrent access with a WaitGroup
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		cache.Set("foo", "bar")
		fmt.Println("Set foo -> bar")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Simulate delay
		if value, found := cache.Get("foo"); found {
			fmt.Printf("Get foo -> %v\n", value)
		} else {
			fmt.Println("foo not found")
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second) // Simulate delay
		cache.Set("baz", "qux")
		fmt.Println("Set baz -> qux")
	}()

	wg.Wait()
}

func main() {
	ThreadSafeCacheExample()
}
