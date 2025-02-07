package main

import (
	"fmt"
	"sync"
)

// SafeMap encapsulates a map and uses a mutex to manage concurrent access.
type SafeMap struct {
	sync.Mutex
	internalMap map[string]int
}

// NewSafeMap initializes a SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		internalMap: make(map[string]int),
	}
}

// Set safely sets a key-value pair in the map.
func (sm *SafeMap) Set(key string, value int) {
	sm.Lock()
	defer sm.Unlock()
	sm.internalMap[key] = value
}

// Get safely retrieves a value from the map by key.
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.Lock()
	defer sm.Unlock()
	value, exists := sm.internalMap[key]
	return value, exists
}

// Delete safely removes a key from the map.
func (sm *SafeMap) Delete(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.internalMap, key)
}

func main() {
	var wg sync.WaitGroup
	safeMap := NewSafeMap()

	// Function to perform concurrent map operations
	worker := func(id int) {
		defer wg.Done()
		key := fmt.Sprintf("key%d", id)

		// Set a value
		safeMap.Set(key, id)
		fmt.Printf("Worker %d set %s to %d\n", id, key, id)

		// Get the value
		if value, exists := safeMap.Get(key); exists {
			fmt.Printf("Worker %d got %d for %s\n", id, value, key)
		}

		// Delete the key
		safeMap.Delete(key)
		fmt.Printf("Worker %d deleted %s\n", id, key)
	}

	// Start multiple goroutines to simulate concurrent access
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("Final map state:")
	for key, value := range safeMap.internalMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
