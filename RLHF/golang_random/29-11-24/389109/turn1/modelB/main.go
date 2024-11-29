package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Simulate lock contention by incrementing a shared counter concurrently
	const numWorkers = 1000
	const numIncrements = 10000
	var counter int

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			for j := 0; j < numIncrements; j++ {
				counter++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final counter value:", counter)

	// Inspect lock contention statistics using the "runtime.ReadMemStats" function
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println("Mutex contention hits:", memStats.MutexContention)
}
