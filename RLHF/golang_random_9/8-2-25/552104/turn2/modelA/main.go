package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem represents a cache item.
type CacheItem struct {
	value      string
	expiration time.Time
}

// PrimaryCache represents a fast in-memory cache.
type PrimaryCache struct {
	sync.RWMutex
	data     map[string]CacheItem
	ttl      time.Duration
	stopChan chan struct{}
}

// SecondaryCache represents a slower storage layer.
type SecondaryCache struct {
	sync.RWMutex
	data map[string]CacheItem
}

// CacheSystem manages both primary and secondary caches.
type CacheSystem struct {
	primary   *PrimaryCache
	secondary *SecondaryCache
}

// NewCacheSystem initializes and returns a new CacheSystem with given TTL for primary cache.
func NewCacheSystem(primaryTTL time.Duration) *CacheSystem {
	pCache := &PrimaryCache{
		data:     make(map[string]CacheItem),
		ttl:      primaryTTL,
		stopChan: make(chan struct{}),
	}

	// Start a goroutine to handle the expiration of primary cache items.
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				pCache.evictExpiredItems()
			case <-pCache.stopChan:
				ticker.Stop()
				return
			}
		}
	}()

	return &CacheSystem{
		primary:   pCache,
		secondary: &SecondaryCache{data: make(map[string]CacheItem)},
	}
}

// Set adds a value to the primary cache.
func (cs *CacheSystem) Set(key, value string) {
	cs.primary.Lock()
	cs.primary.data[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(cs.primary.ttl),
	}
	cs.primary.Unlock()
}

// Get retrieves a value from primary cache, if not found, checks the secondary cache.
func (cs *CacheSystem) Get(key string) (string, bool) {
	// Check primary cache first.
	cs.primary.RLock()
	item, found := cs.primary.data[key]
	cs.primary.RUnlock()
	if found {
		if item.expiration.After(time.Now()) {
			return item.value, true
		}
	}

	// If not found or expired in primary, check secondary.
	cs.secondary.RLock()
	item, found = cs.secondary.data[key]
	cs.secondary.RUnlock()
	if found {
		return item.value, true
	}

	return "", false
}

// Move expired items from primary to secondary cache.
func (p *PrimaryCache) evictExpiredItems() {
	now := time.Now()
	p.Lock()
	defer p.Unlock()
	for key, item := range p.data {
		if now.After(item.expiration) {
			// Move the expired item to the secondary cache.
			fmt.Printf("Moving key %s to secondary cache\n", key)
			delete(p.data, key)

			// Acquire a lock on the secondary cache and add the item.
			p.equalizeToSecondary(key, item)
		}
	}
}

// Transfers item to secondary cache.
func (p *PrimaryCache) equalizeToSecondary(key string, item CacheItem) {
	secondaryCache := &SecondaryCache{
		sync.RWMutex{},
		make(map[string]CacheItem),
	}

	secondaryCache.Lock()
	secondaryCache.data[key] = item
	secondaryCache.Unlock()
}

// Stop expires and stops cleaning process.
func (cs *CacheSystem) Stop() {
	close(cs.primary.stopChan)
}

func main() {
	cacheSystem := NewCacheSystem(5 * time.Second)

	// Set values in the cache
	cacheSystem.Set("key1", "value1")
	cacheSystem.Set("key2", "value2")

	// Sleep for some time to allow TTLs to expire for demonstration
	time.Sleep(6 * time.Second)

	// At this point, items might be in the secondary cache
	fmt.Println("Attempting to retrieve items after TTL expiration.")

	checkAndPrint(cacheSystem, "key1")
	checkAndPrint(cacheSystem, "key2")

	// Cleanup
	cacheSystem.Stop()
}

func checkAndPrint(cs *CacheSystem, key string) {
	value, found := cs.Get(key)
	if found {
		fmt.Printf("Found: %s = %s\n", key, value)
	} else {
		fmt.Printf("%s not found in the cache system\n", key)
	}
}
