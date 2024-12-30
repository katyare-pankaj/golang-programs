package main

import (
	"fmt"
	"sync"
	"time"
)

func addToMap(wg *sync.WaitGroup, m *sync.Map, key int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		m.Store(key, key*i)
	}
}

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	start := time.Now()

	// Create 10 goroutines to add elements to the map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addToMap(&wg, &m, i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Elapsed time: %v\n", elapsed)

	// Print the map to verify
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
