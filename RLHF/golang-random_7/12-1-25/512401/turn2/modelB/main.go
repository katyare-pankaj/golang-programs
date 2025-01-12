package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	sliceSize     = 10000000 // Large slice size to demonstrate memory usage
	numGoroutines = 4        // Number of goroutines for parallel processing
)

func main() {
	// Start memory profiling
	f, err := os.Create("memprofile.prof")
	if err != nil {
		fmt.Println("could not create memory profile:", err)
		return
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile:", err)
	}
	fmt.Println("memory profile written to memprofile.prof")

	// Allocate a large slice
	largeSlice := make([]int, sliceSize)

	// Optionally, you can use a sync.WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Start profiling
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// Process the large slice using goroutines to reduce overhead
	chunkSize := sliceSize / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		go processSlice(largeSlice[start:end], &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Simulate some work after processing the slice
	time.Sleep(5 * time.Second)
}
func processSlice(slice []int, wg *sync.WaitGroup) {
	// Simulate some work on the slice
	for i := range slice {
		slice[i] = i * i
	}

	// Indicate that the goroutine has finished
	wg.Done()
}
