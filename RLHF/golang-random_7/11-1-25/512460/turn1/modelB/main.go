package main

import (
	"fmt"
	"sync"
)

// sharedResource is a resource that multiple goroutines can access.
type sharedResource struct {
	data  map[string]int
	mutex sync.Mutex
}

func (r *sharedResource) increment(key string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[key]++
}

func worker(wg *sync.WaitGroup, r *sharedResource, key string) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		r.increment(key)
	}
}

func main() {
	var wg sync.WaitGroup
	r := &sharedResource{data: make(map[string]int)}

	keys := []string{"key1", "key2", "key3"}

	// Start multiple goroutines that will access the shared resource
	wg.Add(len(keys))
	for _, key := range keys {
		go worker(&wg, r, key)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final values of the shared resource
	for key, value := range r.data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
