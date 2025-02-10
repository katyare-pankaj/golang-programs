package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Generate a large dataset for testing
	dataset := make([]int, 10000000)
	for i := 0; i < len(dataset); i++ {
		dataset[i] = i
	}

	// CPU profiling
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	// Heap profiling
	heapFile, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer heapFile.Close()

	// Process the dataset using range loops to compute the sum of squares
	start := time.Now()
	sum := 0
	for _, value := range dataset {
		sum += value * value
	}
	fmt.Println("Processing time:", time.Since(start))
	fmt.Println("Sum of squares of dataset:", sum)

	// Write heap profile
	if err := pprof.WriteHeapProfile(heapFile); err != nil {
		log.Fatal(err)
	}
}
