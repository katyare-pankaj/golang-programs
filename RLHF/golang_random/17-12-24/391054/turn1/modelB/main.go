package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d starting work...\n", id)
	time.Sleep(time.Second) // Simulate CPU-bound work
	fmt.Printf("Goroutine %d finished work.\n", id)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available cores

	var wg sync.WaitGroup
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	fmt.Printf("Waiting for all goroutines to complete...\n")
	wg.Wait()

	fmt.Printf("All goroutines have completed.\n")
}
