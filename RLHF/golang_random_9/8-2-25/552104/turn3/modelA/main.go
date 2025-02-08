package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem stores the value and its expiration time.
type CacheItem struct {
	value      string
	expiration time.Time
}

// FirstLevelCache is an in-memory cache with a TTL eviction policy.
type FirstLevelCache struct {
	sync.RWMutex
	data     map[string]CacheItem
	ttl      time.Duration
	stopChan chan struct{}
}

// SecondLevelCache simulates persistent storage.
type SecondLevelCache struct {
	sync.RWMutex
	data map[string]string
}

// CacheSystem manages both the first and second level caches.
type CacheSystem struct {
	firstLevel  *FirstLevelCache
	secondLevel *SecondLevelCache
	wg          sync.WaitGroup
}

// NewCacheSystem initializes a new cache system.
func NewCacheSystem(ttl time.Duration) *CacheSystem {
	flCache := &FirstLevelCache{
		data:     make(map[string]CacheItem),
		ttl:      ttl,
		stopChan: make(chan struct{}),
	}

	// Start the background eviction task.
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				flCache.evictExpiredItems()
			case <-flCache.stopChan:
				return
			}
		}
	}()

	return &CacheSystem{
		firstLevel:  flCache,
		secondLevel: &SecondLevelCache{data: make(map[string]string)},
	}
}

// Set adds a new item to the first-level cache and also updates the second-level cache.
func (cs *CacheSystem) Set(key, value string) {
	item := CacheItem{
		value:      value,
		expiration: time.Now().Add(cs.firstLevel.ttl),
	}

	// Update the first-level cache.
	cs.firstLevel.Lock()
	cs.firstLevel.data[key] = item
	cs.firstLevel.Unlock()

	// Update the second-level cache (asynchronous).
	cs.wg.Add(1)
	go func() {
		defer cs.wg.Done()
		cs.secondLevel.Lock()
		cs.secondLevel.data[key] = value
		cs.secondLevel.Unlock()
	}()
}

// Get retrieves an item from the caches, checking first-level before second-level.
func (cs *CacheSystem) Get(key string) (string, bool) {
	// Check in first-level cache.
	cs.firstLevel.RLock()
	item, found := cs.firstLevel.data[key]
	cs.firstLevel.RUnlock()
	if found && item.expiration.After(time.Now()) {
		return item.value, true
	}

	// Check in second-level cache.
	cs.secondLevel.RLock()
	value, found := cs.secondLevel.data[key]
	cs.secondLevel.RUnlock()
	if found {
		// Move the item back to the first-level cache for immediate access next time.
		cs.Set(key, value)
		return value, true
	}

	return "", false
}

// evictExpiredItems removes items from the first-level cache that are expired.
func (fl *FirstLevelCache) evictExpiredItems() {
	now := time.Now()
	fl.Lock()
	for key, item := range fl.data {
		if now.After(item.expiration) {
			delete(fl.data, key)
		}
	}
	fl.Unlock()
}

// Stop initiates the shutdown of background eviction processes.
func (cs *CacheSystem) Stop() {
	close(cs.firstLevel.stopChan)
	cs.wg.Wait()
}

func main() {
	cacheSystem := NewCacheSystem(5 * time.Second)

	// Set values in the cache.
	cacheSystem.Set("key1", "value1")
	cacheSystem.Set("key2", "value2")

	// Simulate some elapsed time.
	time.Sleep(3 * time.Second)

	// Retrieve values to demonstrate persistence.
	fmt.Println("Retrieving soon after inserting:")
	key := "key1"
	if value, found := cacheSystem.Get(key); found {
		fmt.Printf("Found: %s = %s\n", key, value)
	} else {
		fmt.Printf("%s not found\n", key)
	}

	// Allow the first-level cache to expire.
	time.Sleep(3 * time.Second)

	fmt.Println("\nRetrieving after first-level cache expiry:")

	// key1 should have expired in the first-level cache but will persist in the second-level cache.
	if value, found := cacheSystem.Get(key); found {
		fmt.Printf("Found in second-level: %s = %s\n", key, value)
	} else {
		fmt.Printf("%s not found\n", key)
	}

	// Stop the cache eviction process and wait for all updates to complete.
	cacheSystem.Stop()
}
