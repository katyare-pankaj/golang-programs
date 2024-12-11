package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Printf("Atomic Counter: %d, Time: %v\n", counter, time.Since(startTime))
}
