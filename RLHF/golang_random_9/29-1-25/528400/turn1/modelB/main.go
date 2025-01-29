package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	// Sample data slice
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Use sync.WaitGroup to synchronize concurrent processes
	var wg sync.WaitGroup

	// Number of concurrent workers
	numWorkers := 3

	// Start concurrent workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go processData(&wg, data)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("Processing complete.")
}

func processData(wg *sync.WaitGroup, data []int) {
	defer wg.Done()

	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Process data elements
	for _, value := range data {
		fmt.Printf("Processing value: %d\n", value)
	}
}
