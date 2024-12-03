package main

import (
	"fmt"
	"sync"
)

// SafeMap encapsulates a map with a mutex to make it thread-safe.
type SafeMap struct {
	mu    sync.Mutex
	store map[string]int
}

// NewSafeMap initializes a new instance of SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		store: make(map[string]int),
	}
}

// Set adds or updates a key-value pair in the map.
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.store[key] = value
}

// Get retrieves the value associated with the given key.
// It returns the value and a boolean indicating if the key was found.
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.store[key]
	return value, exists
}

// Delete removes a key-value pair from the map.
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.store, key)
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// Number of goroutines for concurrent operations
	numGoroutines := 10

	// Concurrent writing to the map
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	// Wait for all write operations to finish
	wg.Wait()

	// Concurrent reading from the map
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			if value, exists := safeMap.Get(fmt.Sprintf("key%d", i)); exists {
				fmt.Printf("key%d: %d\n", i, value)
			}
		}(i)
	}

	// Wait for all read operations to finish
	wg.Wait()
}
