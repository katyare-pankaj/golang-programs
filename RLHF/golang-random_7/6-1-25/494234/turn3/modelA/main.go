package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sharedData int = 0         // Shared integer variable
var sharedDataMutex sync.Mutex // Mutex to protect sharedData
var sharedDataAtomic int32 = 0 // Atomic integer variable

func incrementData(id int, wg *sync.WaitGroup, useMutex bool, useAtomic bool) {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		if useMutex {
			// Lock the mutex before accessing sharedData
			sharedDataMutex.Lock()
			sharedData++
			sharedDataMutex.Unlock()
		} else if useAtomic {
			// Use atomic operations to safely increment sharedDataAtomic
			atomic.AddInt32(&sharedDataAtomic, 1)
		} else {
			// Unprotected access to sharedData
			sharedData++
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Start goroutines with different synchronization strategies
	wg.Add(1)
	go incrementData(1, &wg, false, false) // No synchronization
	wg.Add(1)
	go incrementData(2, &wg, false, false) // No synchronization

	wg.Add(1)
	go incrementData(3, &wg, true, false) // Using mutex
	wg.Add(1)
	go incrementData(4, &wg, true, false) // Using mutex

	wg.Add(1)
	go incrementData(5, &wg, false, true) // Using atomic
	wg.Add(1)
	go incrementData(6, &wg, false, true) // Using atomic

	wg.Wait()

	fmt.Println("Shared data (no synchronization):", sharedData)
	fmt.Println("Shared data (using mutex):", sharedData)
	fmt.Println("Shared data (using atomic):", int(sharedDataAtomic))
}
