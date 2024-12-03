package main

import (
	"fmt"
	"sync"
	"time"
)

// ThreadSafeMap is a struct that wraps a map and provides thread-safe operations
type ThreadSafeMap[Key, Value any] struct {
	mu   sync.Mutex
	data map[Key]Value
}

// New creates a new thread-safe map
func NewThreadSafeMap[Key comparable, Value any]() *ThreadSafeMap[Key, Value] {
	return &ThreadSafeMap[Key, Value]{data: make(map[Key]Value)}
}

// Get retrieves the value for the given key
func (tsm *ThreadSafeMap[Key, Value]) Get(key Key) (Value, bool) {
	tsm.mu.RLock()
	defer tsm.mu.RUnlock()
	return tsm.data[key]
}

// Set updates or adds the key-value pair
func (tsm *ThreadSafeMap[Key, Value]) Set(key Key, value Value) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()
	tsm.data[key] = value
}

// Delete removes the key-value pair if it exists
func (tsm *ThreadSafeMap[Key, Value]) Delete(key Key) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()
	delete(tsm.data, key)
}

// Range iterates over all key-value pairs
func (tsm *ThreadSafeMap[Key, Value]) Range(f func(key Key, value Value)) {
	tsm.mu.RLock()
	defer tsm.mu.RUnlock()
	for key, value := range tsm.data {
		f(key, value)
	}
}

func main() {
	// Create a thread-safe map
	tsm := NewThreadSafeMap[int, string]()

	// Concurrently update and read the map
	go func() {
		tsm.Set(1, "Value 1")
		tsm.Set(2, "Value 2")
		tsm.Set(3, "Value 3")
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			value, ok := tsm.Get(i)
			if ok {
				fmt.Printf("Key %d: %v\n", i, value)
			} else {
				fmt.Printf("Key %d not found\n", i)
			}
		}
	}()

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)
}
