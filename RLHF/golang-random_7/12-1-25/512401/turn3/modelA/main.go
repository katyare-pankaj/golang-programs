package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	sliceSize     = 10000000 // Size of the slice to process
	numGoroutines = 4        // Number of goroutines to use for processing
)

func main() {
	// Allocate a large slice
	largeSlice := make([]int, sliceSize)

	// Optionally, you can use a sync.WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Start profiling
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

	// Start CPU profiling
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// Process the large slice using goroutines
	chunkSize := sliceSize / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := min(start+chunkSize, sliceSize)
		go processSlice(largeSlice[start:end], &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Simulate some additional work after processing the slice
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
