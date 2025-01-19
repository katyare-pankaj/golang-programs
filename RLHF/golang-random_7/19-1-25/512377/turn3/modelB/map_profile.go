package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

func processData(input []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// some processing on input data
	mapData := make(map[int]int)
	for _, value := range input {
		mapData[value]++
	}
	// Replace the map data with any other data structure if needed for faster access
	mapData = make(map[int]int) // Freeing the previously allocated memory (optional)
}

func main() {
	// Disable GC for a more precise profile
	runtime.GC()

	// Generate a large dataset for testing
	const datasetSize = 10_000_000
	dataset := make([]int, datasetSize)
	for i := 0; i < datasetSize; i++ {
		dataset[i] = i
	}

	numThreads := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	wg.Add(numThreads)

	// Start memory profile
	mProf := pprof.Lookup("heap")
	if mProf == nil {
		fmt.Println("Memory profiler not available.")
		return
	}

	// Calculate memory usage before processing data
	if err := mProf.WriteTo(os.Stdout, 0); err != nil {
		fmt.Println("Error writing memory profile:", err)
	}

	partSize := datasetSize / numThreads
	for i := 0; i < numThreads; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		go processData(dataset[start:end], wg)
	}

	wg.Wait()

	// Calculate memory usage after processing data
	if err := mProf.WriteTo(os.Stdout, 0); err != nil {
		fmt.Println("Error writing memory profile:", err)
	}
}
