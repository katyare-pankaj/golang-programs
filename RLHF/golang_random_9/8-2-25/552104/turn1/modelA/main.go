package main

import (
	"fmt"
	"sync"
	"time"
)

// FirstLevelCache represents an in-memory cache.
type FirstLevelCache struct {
	sync.RWMutex
	data map[string]string
}

// SecondLevelCache simulates a slower persistent storage.
type SecondLevelCache struct {
	sync.RWMutex
	data map[string]string
}

// CacheSystem encapsulates both cache levels.
type CacheSystem struct {
	firstLevel  *FirstLevelCache
	secondLevel *SecondLevelCache
	wg          sync.WaitGroup
}

// NewCacheSystem initializes a new multi-level cache system.
func NewCacheSystem() *CacheSystem {
	return &CacheSystem{
		firstLevel:  &FirstLevelCache{data: make(map[string]string)},
		secondLevel: &SecondLevelCache{data: make(map[string]string)},
	}
}

// Get retrieves a value from the cache system.
func (cs *CacheSystem) Get(key string) (string, bool) {
	// Try to get from first-level cache
	cs.firstLevel.RLock()
	value, ok := cs.firstLevel.data[key]
	cs.firstLevel.RUnlock()
	if ok {
		return value, true
	}

	// If not found, try to get from second-level cache
	cs.secondLevel.RLock()
	value, ok = cs.secondLevel.data[key]
	cs.secondLevel.RUnlock()
	if ok {
		// Update first-level cache asynchronously
		cs.wg.Add(1)
		go func() {
			defer cs.wg.Done()
			cs.firstLevel.Lock()
			cs.firstLevel.data[key] = value
			cs.firstLevel.Unlock()
		}()
		return value, true
	}

	return "", false
}

// Set adds a value to the cache system.
func (cs *CacheSystem) Set(key, value string) {
	// Update both caches concurrently
	cs.wg.Add(2)
	go func() {
		defer cs.wg.Done()
		cs.firstLevel.Lock()
		cs.firstLevel.data[key] = value
		cs.firstLevel.Unlock()
	}()

	go func() {
		defer cs.wg.Done()
		time.Sleep(100 * time.Millisecond) // Simulate a slower write operation
		cs.secondLevel.Lock()
		cs.secondLevel.data[key] = value
		cs.secondLevel.Unlock()
	}()
}

func main() {
	cacheSystem := NewCacheSystem()

	// Set some values in the cache
	cacheSystem.Set("key1", "value1")
	cacheSystem.Set("key2", "value2")

	// Wait for all cache updates to finish
	cacheSystem.wg.Wait()

	// Retrieve values
	if value, found := cacheSystem.Get("key1"); found {
		fmt.Printf("Found: key1 = %s\n", value)
	} else {
		fmt.Println("key1 not found")
	}

	if value, found := cacheSystem.Get("key3"); found {
		fmt.Printf("Found: key3 = %s\n", value)
	} else {
		fmt.Println("key3 not found")
	}
}
