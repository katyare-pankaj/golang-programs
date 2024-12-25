package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Create a file to write the CPU profile
	cpuProfileFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile file:", err)
		return
	}
	defer cpuProfileFile.Close()

	// Start CPU profiling
	err = pprof.StartCPUProfile(cpuProfileFile)
	if err != nil {
		fmt.Println("Error starting CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Record the start time
	start := time.Now()
	defer fmt.Println("Total time:", time.Since(start))

	// Run the benchmark function multiple times to get an average
	const numBenchmarks = 1000000
	for i := 0; i < numBenchmarks; i++ {
		benchmarkDefer()
	}
}

func benchmarkDefer() {
	// Simulate some work being done
	for i := 0; i < 100; i++ {
		_ = i * i
	}

	// Simulate resource cleanup with defer
	defer func() {
		time.Sleep(time.Nanosecond * 1)
	}()
}
