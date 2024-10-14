// Simple in-memory cache using a map
package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	value  interface{}
	expiry time.Time
}
type Cache struct {
	mu            sync.RWMutex
	items         map[string]CacheItem
	defaultExpiry time.Duration
}

func NewCache(defaultExpiry time.Duration) *Cache {
	return &Cache{
		items:         make(map[string]CacheItem),
		defaultExpiry: defaultExpiry,
	}
}
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item := CacheItem{value: value, expiry: time.Now().Add(c.defaultExpiry)}
	c.items[key] = item
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok || item.expiry.Before(time.Now()) {
		return nil, false
	}
	return item.value, true
}
func main() {
	cache := NewCache(5 * time.Minute) // Cache items for 5 minutes
	cache.Set("key1", "value1")
	if val, ok := cache.Get("key1"); ok {
		fmt.Println("Cache Hit:", val) // Output: Cache Hit: value1
	}
	time.Sleep(6 * time.Minute)
	if _, ok := cache.Get("key1"); !ok {
		fmt.Println("Cache Miss") // Output: Cache Miss
	}
}
