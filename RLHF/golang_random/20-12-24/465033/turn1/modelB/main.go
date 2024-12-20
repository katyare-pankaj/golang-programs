package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 100
	numTasks   = 100000
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < numTasks; i++ {
		// Simulate work by doing some computations.
		_ = id * i
	}
}

func main() {
	// With synchronization.
	wg := sync.WaitGroup{}
	start := time.Now()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, i)
	}
	wg.Wait()
	elapsedWithSync := time.Since(start)
	fmt.Printf("Elapsed time with synchronization: %s\n", elapsedWithSync)

	// Without synchronization.
	start = time.Now()
	for i := 0; i < numWorkers; i++ {
		go worker(nil, i)
	}
	// Wait for all workers to finish without using WaitGroup.
	time.Sleep(time.Duration(numWorkers) * time.Second)
	elapsedWithoutSync := time.Since(start)
	fmt.Printf("Elapsed time without synchronization: %s\n", elapsedWithoutSync)

	// Calculate the overhead.
	overhead := elapsedWithoutSync - elapsedWithSync
	fmt.Printf("Performance overhead with synchronization: %s\n", overhead)
}
