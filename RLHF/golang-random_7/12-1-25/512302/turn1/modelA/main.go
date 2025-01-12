package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
)

// Simple data structure to hold key-value pairs
type DataEntry struct {
	Key   string
	Value string
}

// In-memory cache implemented as a map
type InMemoryCache struct {
	cache map[string]DataEntry
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{cache: make(map[string]DataEntry)}
}

func (imc *InMemoryCache) Get(key string) (DataEntry, bool) {
	return imc.cache[key], imc.cache[key] != (DataEntry{})
}

func (imc *InMemoryCache) Set(entry DataEntry) {
	imc.cache[entry.Key] = entry
}

// Disk cache implementation
type DiskCache struct {
	filePath string
}

func NewDiskCache(filePath string) *DiskCache {
	return &DiskCache{filePath: filePath}
}

func (dc *DiskCache) Get(key string) (DataEntry, bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error retrieving data from disk: %v", err)
		}
	}()

	data, err := ioutil.ReadFile(dc.filePath)
	if err != nil || len(data) == 0 {
		return DataEntry{}, false
	}

	var cache map[string]DataEntry
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&cache); err != nil {
		log.Printf("Error decoding cache from disk: %v", err)
		return DataEntry{}, false
	}

	entry, exists := cache[key]
	return entry, exists
}

func (dc *DiskCache) Set(entry DataEntry) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error writing data to disk: %v", err)
		}
	}()

	cache := make(map[string]DataEntry)

	data, err := ioutil.ReadFile(dc.filePath)
	if err == nil {
		var existingCache map[string]DataEntry
		if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&existingCache); err == nil {
			cache = existingCache
		}
	}

	cache[entry.Key] = entry

	buffer := new(bytes.Buffer)
	if err := gob.NewEncoder(buffer).Encode(cache); err != nil {
		log.Printf("Error encoding cache to disk: %v", err)
		return
	}

	if err := ioutil.WriteFile(dc.filePath, buffer.Bytes(), 0644); err != nil {
		log.Printf("Error writing cache to disk: %v", err)
	}
}

func main() {
	inMemoryCache := NewInMemoryCache()
	diskCache := NewDiskCache("disk_cache.gob")

	// Sample data for demonstration
	dataEntries := []DataEntry{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}

	// Initialize the disk cache with data
	for _, entry := range dataEntries {
		diskCache.Set(entry)
	}

	// Load the data into the in-memory cache
	for _, entry := range dataEntries {
		inMemoryCache.Set(entry)
	}

	// Range loop to iterate over keys and retrieve data using the caches
	fmt.Println("Retrieving data with range loop and multi-level caching:")
	for _, key := range []string{"key1", "key2", "key3", "key4"} { // include a non-existing key for demonstration
		entry, exists := inMemoryCache.Get(key)
		if exists {
			fmt.Printf("Found key '%s' in in-memory cache: %s\n", key, entry.Value)
		} else {
			entry, exists = diskCache.Get(key)
			if exists {
				fmt.Printf("Found key '%s' in disk cache: %s\n", key, entry.Value)
				// Load the found entry into the in-memory cache
				inMemoryCache.Set(entry)
			} else {
				fmt.Printf("Key '%s' not found in either cache.\n", key)
			}
		}
	}
}
