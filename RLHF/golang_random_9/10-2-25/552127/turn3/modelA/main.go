package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Create a large dataset for testing
	dataset := make([]int, 10000000)
	for i := 0; i < len(dataset); i++ {
		dataset[i] = i + 1
	}

	// Start CPU profiling
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile: ", err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("Could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Compute the sum of squares
	start := time.Now()
	sumOfSquares := 0
	for _, value := range dataset {
		sumOfSquares += value * value
	}
	fmt.Println("Processing time:", time.Since(start))
	fmt.Println("Sum of squares:", sumOfSquares)

	// Write heap profile
	heapFile, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal("Could not create heap profile: ", err)
	}
	defer heapFile.Close()

	if err := pprof.WriteHeapProfile(heapFile); err != nil {
		log.Fatal("Could not write heap profile: ", err)
	}
}