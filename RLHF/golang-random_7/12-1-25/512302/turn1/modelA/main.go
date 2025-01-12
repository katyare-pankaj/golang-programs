package main

import (
	"fmt"
	"time"
)

type cache struct {
	data map[int]string
}

func newCache() *cache {
	return &cache{
		data: make(map[int]string),
	}
}

func (c *cache) get(key int) (string, bool) {
	return c.data[key], c.data[key] != ""
}

func (c *cache) set(key int, value string) {
	c.data[key] = value
}

func fetchFromExternal(key int) string {
	// Simulate external data fetching by adding a delay
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("External data for key %d", key)
}

func main() {
	// Initialize caches
	l1Cache := newCache()
	l2Cache := newCache()
	l3Cache := newCache()

	// Simulated data range
	for key := 0; key < 10; key++ {
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
