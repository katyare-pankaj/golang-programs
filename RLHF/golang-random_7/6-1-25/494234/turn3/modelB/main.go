package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

// Counter is a simple struct to demonstrate shared data
type Counter struct {
	value int
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.value++
}

// Dec decrements the counter
func (c *Counter) Dec() {
	c.value--
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Number of goroutines accessing the shared counter
	const numGoroutines = 100

	// Start the goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// Simulate random access to the counter
				if rand.Intn(2) == 0 {
					counter.Inc()
				} else {
					counter.Dec()
				}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Unsafe access to the shared counter
	fmt.Println("Unsafe Counter Value:", counter.value)

	// Mitigate concurrency issues using Mutex
	var mutex sync.Mutex
	var safeCounter Counter

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// Simulate random access to the counter
				if rand.Intn(2) == 0 {
					mutex.Lock()
					safeCounter.Inc()
					mutex.Unlock()
				} else {
					mutex.Lock()
					safeCounter.Dec()
					mutex.Unlock()
				}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Safe access to the shared counter using Mutex
	fmt.Println("Safe Counter Value using Mutex:", safeCounter.value)

	// Mitigate concurrency issues using Atomic Operations
	var atomicCounter int32

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// Simulate random access to the counter
				if rand.Intn(2) == 0 {
					// Increment the atomic counter using atomic.AddInt32
					atomic.AddInt32(&atomicCounter, 1)
				} else {
					// Decrement the atomic counter using atomic.AddInt32
					atomic.AddInt32(&atomicCounter, -1)
				}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Safe access to the shared counter using Atomic Operations
	fmt.Println("Safe Counter Value using Atomic Operations:", atomicCounter)
}
