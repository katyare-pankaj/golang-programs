package main  
import (  
    "fmt"
    "sync"
    "time"
)
// CacheItem represents an item in the cache
type CacheItem struct {  
    value interface{}
    expireAt time.Time
}
// MultilevelCache is a multilevel cache with LRU and TTL
type MultilevelCache struct {  
    level1 map[string]*CacheItem
    level2 map[string]*CacheItem
    lock sync.Mutex
    wg sync.WaitGroup
    ttl time.Duration
}
// NewMultilevelCache creates a new multilevel cache
func NewMultilevelCache(ttl time.Duration) *MultilevelCache {  
    return &MultilevelCache{  
        level1: make(map[string]*CacheItem),
        level2: make(map[string]*CacheItem),
        ttl: ttl,
    }
}
// getFromCache retrieves an item from the cache if it exists and is not expired
func (c *MultilevelCache) getFromCache(key string) (interface{}, bool) {  
    c.lock.Lock()
    defer c.lock.Unlock()

    // Check if the item exists in level 1 cache
    if item, ok := c.level1[key]; ok && time.Now().Before(item.expireAt) {
        return item.value, true
    }

    // Check if the item exists in level 2 cache
    if item, ok := c.level2[key]; ok && time.Now().Before(item.expireAt) {
        // Move the item to level 1 cache
        c.level1[key] = item
        delete(c.level2, key)
        return item.value, true
    }
    return nil, false
}
// setToCache adds or updates an item in the cache
func (c *MultilevelCache) setToCache(key string, value interface{}) {  
    c.lock.Lock()
    defer c.lock.Unlock()

    expireAt := time.Now().Add(c.ttl)
    item := &CacheItem{value: value, expireAt: expireAt}
    c.level1[key] = item

    // Start a goroutine to move the item to level 2 cache after TTL
    c.wg.Add(1)
    go func() {  
        time.Sleep(c.ttl)
        c.lock.Lock()
        defer c.lock.Unlock()

        // If the item is still in level 1 cache, move it to level 2 cache
        if _, ok := c.level1[key]; ok {
            c.level2[key] = item
            delete(c.level1, key)
        }
        c.wg.Done()
    }()
}
// WaitForPendingUpdates waits for all pending cache updates to complete
func (c *MultilevelCache) WaitForPendingUpdates() {  
    c.wg.Wait()
}

func main() {  
    cache := NewMultilevelCache(5 * time.Second)

    // Add items to the cache
    cache.setToCache("key1", "value1")
    cache.setToCache("key2", "value2")

    // Get items from the cache
    value1, ok := cache.getFromCache("key1")
    if ok {
        fmt.Println("key1:", value1)
    } else {
        fmt.Println("key1 not found")
    }

    value2, ok := cache.getFromCache("key2")
    if ok {
        fmt.Println("key2:", value2)
    } else {
        fmt.Println("key2 not found")
    }
    // Wait for all pending updates to complete
    cache.WaitForPendingUpdates()

    // After TTL, items should be moved to level 2 cache
    value1, ok = cache.getFromCache("key1")
    if ok {
        fmt.Println("key1:", value1)
    } else {
        fmt.Println("key1 not found")
    }

    value2, ok = cache.getFromCache("key2")
    if ok {