package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Generate a large dataset for testing
	dataset := make([]int, 1000000)
	for i := 0; i < len(dataset); i++ {
		dataset[i] = i
	}

	// Create file for CPU profiling
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("could not create CPU profile: %v", err)
	}
	defer cpuFile.Close() // ensure that we close the file

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatalf("could not start CPU profile: %v", err)
	}
	defer pprof.StopCPUProfile() // ensure that we stop profiling

	// Process the dataset using range loops
	start := time.Now()
	sum := processDataset(dataset)
	fmt.Println("Processing time:", time.Since(start))
	fmt.Println("Sum of dataset:", sum)

	// Create file for heap profiling
	heapFile, err := os.Create("heap.prof")
	if err != nil {
		log.Fatalf("could not create heap profile: %v", err)
	}
	defer heapFile.Close() // ensure that we close the file

	// Force garbage collection before writing heap profile
	runtime.GC()

	// Generate heap profile
	if err := pprof.WriteHeapProfile(heapFile); err != nil {
		log.Fatalf("could not write heap profile: %v", err)
	}
}

func processDataset(dataset []int) int {
	sum := 0
	for _, value := range dataset {
		sum += value
	}
	return sum
}
