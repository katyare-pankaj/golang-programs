package main

import (
	"fmt"
	"time"
)

type cache struct {
	data map[int]string
	max  int
}

func newCache(max int) *cache {
	return &cache{
		data: make(map[int]string),
		max:  max,
	}
}

func (c *cache) get(key int) (string, bool) {
	value, found := c.data[key]
	if found {
		// Simulate cache hit by moving to front
		delete(c.data, key)
		c.data[key] = value
	}
	return value, found
}

func (c *cache) set(key int, value string) {
	if len(c.data) >= c.max {
		// Evict the least recently used key
		for keyToEvict := range c.data {
			delete(c.data, keyToEvict)
			break
		}
	}
	c.data[key] = value
}

func fetchFromExternal(key int) string {
	// Simulate external data fetching by adding a delay
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("External data for key %d", key)
}

func main() {
	// Initialize caches with different sizes
	l1Cache := newCache(10)
	l2Cache := newCache(50)
	l3Cache := newCache(100)

	// Simulated data range
	for key := 0; key < 150; key++ {
		value, found := l1Cache.get(key)
		if found {
			fmt.Printf("Key %d found in L1 cache: %s\n", key, value)
		} else {
			value, found = l2Cache.get(key)
			if found {
				fmt.Printf("Key %d found in L2 cache: %s\n", key, value)
				l1Cache.set(key, value) // Update L1 cache with value from L2
			} else {
				value, found = l3Cache.get(key)
				if found {
					fmt.Printf("Key %d found in L3 cache: %s\n", key, value)
					l2Cache.set(key, value) // Update L2 cache with value from L3
					l1Cache.set(key, value) // Update L1 cache with value from L3
				} else {
					// Fetch from external source if not found in all caches
					value = fetchFromExternal(key)
					fmt.Printf("Key %d fetched from external source: %s\n", key, value)
					l3Cache.set(key, value) // Update L3 cache with fetched value
					l2Cache.set(key, value) // Update L2 cache with fetched value
					l1Cache.set(key, value) // Update L1 cache with fetched value
				}
			}
		}
	}
}
