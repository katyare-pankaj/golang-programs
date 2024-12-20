package main

import (
	"fmt"
	"sync"
	"time"
)

const workers = 1000
const iterations = 100000

func main() {
	// Measure without WaitGroup
	start := time.Now()
	executeWithoutWaitGroup()
	fmt.Printf("Execution time without WaitGroup: %v\n", time.Since(start))

	// Measure with WaitGroup
	start = time.Now()
	executeWithWaitGroup()
	fmt.Printf("Execution time with WaitGroup: %v\n", time.Since(start))
}

func executeWithoutWaitGroup() {
	for i := 0; i < workers; i++ {
		go work()
	}
	// Sleep to wait for goroutines to finish
	time.Sleep(time.Millisecond * 500)
}

func executeWithWaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}

func work() {
	for j := 0; j < iterations; j++ {
		_ = j * j
	}
}
