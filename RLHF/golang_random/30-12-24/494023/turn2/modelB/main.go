package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// Define a set of predefined keys
	keys := []string{"key1", "key2", "key3", "key4", "key5"}

	// Create a sync.Map to store counters for each key
	var counters sync.Map

	// Number of goroutines to increment the counters
	const numGoroutines = 10

	// Create a WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Start the goroutines to increment the counters concurrently
	for i := 0; i < numGoroutines; i++ {
		go incrementCounters(&counters, keys, &wg)
	}

	// Wait for all goroutines to complete their execution
	wg.Wait()

	// Display the final counts for each key
	counters.Range(func(key, value interface{}) bool {
		fmt.Printf("Count for %s: %d\n", key.(string), value.(int))
		return true
	})
}

func incrementCounters(counters *sync.Map, keys []string, wg *sync.WaitGroup) {
	// Defer the WaitGroup Done function to signal that this goroutine is complete
	defer wg.Done()

	// Loop for a random number of times to increment the counters
	for i := 0; i < rand.Intn(100); i++ {
		// Select a random key from the set
		key := keys[rand.Intn(len(keys))]

		// Load the current count for the selected key using LoadOrStore
		count, _ := counters.LoadOrStore(key, 0)

		// Increment the count and store it back using Store
		counters.Store(key, count.(int)+1)
	}
}
