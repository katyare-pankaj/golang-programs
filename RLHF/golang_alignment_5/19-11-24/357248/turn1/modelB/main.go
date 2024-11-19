package main

import (
	"fmt"
	"sync"
)

const (
	goroutineCount = 10000
	incrementCount = 1000
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	// Increment the counter concurrently from multiple goroutines
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementCount; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
