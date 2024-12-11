package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mutex         sync.Mutex
	atomicCounter int64 = 0
	mutexCounter  int64 = 0
	wg            sync.WaitGroup
)

func incrementWithMutex() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		mutexCounter++
		mutex.Unlock()
	}
}

func incrementAtomic() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(&atomicCounter, 1)
	}
}

func main() {
	startTime := time.Now()

	numGoroutines := 100
	wg.Add(numGoroutines * 2)

	// Increment with Mutex
	for i := 0; i < numGoroutines; i++ {
		go incrementWithMutex()
	}

	// Increment with Atomic
	for i := 0; i < numGoroutines; i++ {
		go incrementAtomic()
	}

	wg.Wait()

	fmt.Printf("Mutex Counter: %d, Time: %v\n", mutexCounter, time.Since(startTime))
	fmt.Printf("Atomic Counter: %d, Time: %v\n", atomicCounter, time.Since(startTime))
}
