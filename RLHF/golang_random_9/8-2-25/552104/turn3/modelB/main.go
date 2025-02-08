package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem represents a cached item.
type CacheItem struct {
	value      string
	expiration time.Time
}

// FirstLevelCache represents a TTL-based in-memory cache.
type FirstLevelCache struct {
	sync.RWMutex
	data     map[string]CacheItem
	ttl      time.Duration
	stopChan chan struct{}
}

// NewFirstLevelCache creates a new FirstLevelCache with a given TTL.
func NewFirstLevelCache(ttl time.Duration) *FirstLevelCache {
	return &FirstLevelCache{
		data:     make(map[string]CacheItem),
		ttl:      ttl,
		stopChan: make(chan struct{}),
	}
}

// StartEviction starts a goroutine to evict expired items from the cache.
func (f *FirstLevelCache) StartEviction() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				f.evictExpiredItems()
			case <-f.stopChan:
				ticker.Stop()
				return
			}
		}
	}()
}

// StopEviction stops the eviction process.
func (f *FirstLevelCache) StopEviction() {
	close(f.stopChan)
}

// Set adds a value to the cache with an expiration time.
func (f *FirstLevelCache) Set(key, value string) {
	f.Lock()
	defer f.Unlock()
	f.data[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(f.ttl),
	}
}

// Get retrieves a value from the cache if not expired.
func (f *FirstLevelCache) Get(key string) (string, bool) {
	f.RLock()
	defer f.RUnlock()
	item, found := f.data[key]
	if found && item.expiration.After(time.Now()) {
		return item.value, true
	}
	return "", false
}

// evictExpiredItems removes expired items from the cache.
func (f *FirstLevelCache) evictExpiredItems() {
	now := time.Now()
	f.Lock()
	defer f.Unlock()
	for key, item := range f.data {
		if now.After(item.expiration) {
			delete(f.data, key)
		}
	}
}

// SecondLevelCache represents a persistent in-memory cache.
type SecondLevelCache struct {
	sync.RWMutex
	data map[string]string
}

// NewSecondLevelCache creates a new SecondLevelCache.
func NewSecondLevelCache() *SecondLevelCache {
	return &SecondLevelCache{
		data: make(map[string]string),
	}
}

// Set adds a value to the second-level cache.
func (s *SecondLevelCache) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

// Get retrieves a value from the second-level cache.
func (s *SecondLevelCache) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	value, found := s.data[key]
	return value, found
}

// CacheSystem manages both first and second-level caches.
type CacheSystem struct {
	firstLevel  *FirstLevelCache
	secondLevel *SecondLevelCache
}

// NewCacheSystem creates a new CacheSystem with a given TTL for the first-level cache.
func NewCacheSystem(ttl time.Duration) *CacheSystem {
	return &CacheSystem{
		firstLevel:  NewFirstLevelCache(ttl),
		secondLevel: NewSecondLevelCache(),
	}
}

// Start starts the caching system, including eviction for the first-level cache.
func (cs *CacheSystem) Start() {
	cs.firstLevel.StartEviction()
}

// Stop stops the caching system, including eviction for the first-level cache.
func (cs *CacheSystem) Stop() {
	cs.firstLevel.StopEviction()
}

// Set adds a value to the cache system, first to the first-level cache with TTL, then to the second-level cache.
func (cs *CacheSystem) Set(key, value string) {
	cs.firstLevel.Set(key, value)
	cs.secondLevel.Set(key, value)
}

// Get retrieves a value from the cache system, first from the first-level cache, then from the second-level cache.
func (cs *CacheSystem) Get(key string) (string, bool) {
	value, found := cs.firstLevel.Get(key)
	if found {
		return value, true
	}
	return cs.secondLevel.Get(key)
}
func main() {
	cacheSystem := NewCacheSystem(2 * time.Second)
	cacheSystem.Start()
	defer cacheSystem.Stop()
	// Set some values in the cache
	cacheSystem.Set("key1", "value1")
	cacheSystem.Set("key2", "value2")
	// Sleep for some time to allow TTL to expire for demonstration
	time.Sleep(3 * time.Second)
	fmt.Println("Attempting to retrieve items after TTL expiration.")
	checkAndPrint(cacheSystem, "key1")
	checkAndPrint(cacheSystem, "key2")
}
func checkAndPrint(cs *CacheSystem, key string) {
	value, found := cs.Get(key)
	if found {
		fmt.Printf("Found: %s = %s\n", key, value)
	} else {
		fmt.Printf("%s not found in the cache system\n", key)
	}
}
