package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Key   int
	Value string
}

func main() {
	// Simulate some data
	data := []Data{
		{1, "value1"},
		{2, "value2"},
		{3, "value3"},
		{4, "value4"},
		{5, "value5"},
	}

	// Create multi-level caches
	level1Cache := make(map[int]string)
	level2Cache := make(map[int]*sync.Mutex)

	var wg sync.WaitGroup
	wg.Add(len(data))

	// Go routine to populate level 1 cache with initial data
	go func() {
		for _, d := range data {
			level1Cache[d.Key] = d.Value
			level2Cache[d.Key] = new(sync.Mutex)
			wg.Done()
		}
	}()
	wg.Wait()

	fmt.Println("Level 1 Cache:", level1Cache)

	// Function to retrieve data from multi-level cache
	getData := func(key int) string {
		// First check the level 1 cache
		if value, ok := level1Cache[key]; ok {
			return value
		}

		// If not in level 1, check level 2
		mutex, ok := level2Cache[key]
		if !ok {
			return "" // Key not found
		}

		// Lock the level 2 cache for this key to avoid race conditions
		mutex.Lock()
		defer mutex.Unlock()

		// Check level 2 cache again after acquiring the lock
		if value, ok := level1Cache[key]; ok {
			return value
		}

		// Simulate retrieving data from a slow source (e.g., database)
		// In this example, we'll just simulate it by adding a delay
		// For demonstration purposes only.
		value := fmt.Sprintf("value-from-slow-source-for-key-%d", key)

		// Update level 1 cache
		level1Cache[key] = value

		return value
	}

	// Use the multi-level cache in a range loop
	for _, key := range []int{1, 2, 3, 4, 5, 6, 7} {
		// Retrieve data using the getData function
		value, _ := getData(key)
		fmt.Println("Data for key", key, ":", value)
	}
}
