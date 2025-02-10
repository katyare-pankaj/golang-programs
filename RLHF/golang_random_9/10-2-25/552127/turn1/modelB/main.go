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
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// Process the dataset using range loops
	start := time.Now()
	sum := 0
	for _, value := range dataset {
		sum += value
	}
	fmt.Println("Processing time:", time.Since(start))
	fmt.Println("Sum of dataset:", sum)
	// Force garbage collection to get a clean profile
	runtime.GC()
	// Generate heap profile
	heapFile, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(heapFile)
	heapFile.Close()
}
