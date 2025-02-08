package main  
import (  
    "fmt"
    "sync"
    "time"
)

type Item struct {  
    value string
    expiry time.Time
}

// FirstLevelCache represents a time-to-live (TTL) cache with a fixed size.
type FirstLevelCache struct {  
    items map[string]Item
    maxSize int
    mu sync.RWMutex
    ttl time.Duration
}

func NewFirstLevelCache(maxSize int, ttl time.Duration) *FirstLevelCache {  
    return &FirstLevelCache{
        items: make(map[string]Item),
        maxSize: maxSize,
        ttl: ttl,
    }
}

// Get retrieves an item from the first-level cache.
func (c *FirstLevelCache) Get(key string) (string, bool) {  
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, ok := c.items[key]
    if !ok || time.Now().After(item.expiry) {
        return "", false
    }
    
    // Refresh the TTL of the accessed item
    c.items[key].expiry = time.Now().Add(c.ttl)
    return item.value, true
}

// Set adds or updates an item in the first-level cache.
func (c *FirstLevelCache) Set(key, value string) {  
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = Item{value: value, expiry: time.Now().Add(c.ttl)}

    // Evict least recently used item if cache is full
    if len(c.items) > c.maxSize {
        // Find the oldest item
        var oldestKey string
        var oldestExpiry time.Time
        for key, item := range c.items {
            if oldestExpiry.IsZero() || item.expiry.Before(oldestExpiry) {
                oldestKey = key
                oldestExpiry = item.expiry
            }
        }
        delete(c.items, oldestKey)
    }
}

// SecondLevelCache represents a simple, persistent cache.
type SecondLevelCache struct {  
    data map[string]string
    mu sync.RWMutex
}

func NewSecondLevelCache() *SecondLevelCache {  
    return &SecondLevelCache{
        data: make(map[string]string),
    }
}

// Get retrieves a value from the second-level cache.
func (c *SecondLevelCache) Get(key string) (string, bool) {  
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, ok := c.data[key]
    return value, ok
}

// Set adds a value to the second-level cache.
func (c *SecondLevelCache) Set(key, value string) {  
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

// CacheSystem encapsulates both cache levels.
type CacheSystem struct {  
    firstLevel *FirstLevelCache
    secondLevel *SecondLevelCache
}

// NewCacheSystem initializes a new multi-level cache system.
func NewCacheSystem(firstLevelSize int, firstLevelTTL time.Duration) *CacheSystem {  
    return &CacheSystem{
        firstLevel:  NewFirstLevelCache(firstLevelSize, firstLevelTTL),
        secondLevel: NewSecondLevelCache(),
    }
}

// Get retrieves a value from the cache system.
func (cs *CacheSystem) Get(key string) (string, bool) {  
    // Try to get from first-level cache
    value, ok := cs.firstLevel.Get(key)
    if ok {
        return value, true
    }

    // If not found, try to get from second-level cache
    value, ok = cs.secondLevel.Get(key)
    if ok {
        // Move item to first-level cache