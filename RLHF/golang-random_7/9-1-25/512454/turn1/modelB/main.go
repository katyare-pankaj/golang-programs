package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Collect Goroutine metrics and identify imbalances
func main() {
	const numTasks = 1000
	const maxGoroutines = 10
	taskChan := make(chan int, numTasks)

	// Limit the number of active Goroutines using a semaphore
	sem := make(chan struct{}, maxGoroutines)
	defer close(sem)

	// Count the active Goroutines and generate tasks
	var wg sync.WaitGroup
	var activeGoroutines int

	wg.Add(numTasks)
	go func() {
		defer wg.Done()
		for i := 0; i < numTasks; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	// Worker Goroutines that consume tasks
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChan {
				// Simulate work
				time.Sleep(time.Millisecond * 10)

				// Check for imbalances in task distribution
				activeGoroutines++
				if activeGoroutines > maxGoroutines {
					fmt.Println("WARNING: Imbalance in task distribution!")
				}
				activeGoroutines--

				sem <- struct{}{}        // Acquire a semaphore to limit Goroutines
				defer func() { <-sem }() // Release the semaphore after work
			}
		}()
	}

	wg.Wait()

	// Collect and print runtime metrics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc: %v MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys: %v MB\n", m.Sys/1024/1024)
	fmt.Printf("Lookups: %v\n", m.Lookups)
	fmt.Printf("Mallocs: %v\n", m.Mallocs)
	fmt.Printf("Frees: %v\n", m.Frees)
	fmt.Printf("HeapAlloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("HeapSys: %v MB\n", m.HeapSys/1024/1024)
	fmt.Printf("HeapIdle: %v MB\n", m.HeapIdle/1024/1024)
	fmt.Printf("HeapInuse: %v MB\n", m.HeapInuse/1024/1024)
	fmt.Printf("HeapReleased: %v MB\n", m.HeapReleased/1024/1024)
	fmt.Printf("HeapObjects: %v\n", m.HeapObjects)
}
