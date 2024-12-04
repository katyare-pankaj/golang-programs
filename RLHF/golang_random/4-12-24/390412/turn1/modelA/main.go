package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var Counter int64 = 0

func incrementCounter() {
	atomic.AddInt64(&Counter, 1)
}

func main() {
	const numGoroutines = 100
	const operationsPerGoroutine = 100000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < operationsPerGoroutine; j++ {
				incrementCounter()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	endTime := time.Now()

	fmt.Printf("Execution time: %.2f seconds\n", endTime.Sub(startTime).Seconds())
	fmt.Printf("Final counter value: %d\n", Counter)
}
