package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sharedCounter uint64

func incrementCounterAtomic() {
	for i := 0; i < 10000; i++ {
		atomic.AddUint64(&sharedCounter, 1)
	}
}

func main() {
	// Create multiple goroutines to increment the counter
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			incrementCounterAtomic()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value with atomic:", sharedCounter)
}
