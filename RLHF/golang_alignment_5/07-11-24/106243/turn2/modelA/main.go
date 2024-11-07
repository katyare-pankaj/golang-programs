package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	// Enable memory profiling
	runtime.MemProfileRate = 1024 // Sample every 1KB allocation

	// Simulate memory allocation
	doSomeMemoryWork()

	// Take a memory profile snapshot
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("could not create memory profile: ", err)
		return
	}
	defer f.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile: ", err)
	}

	fmt.Println("Memory profiling completed.")
}

func doSomeMemoryWork() {
	var data []int
	for i := 0; i < 1000000; i++ {
		// Append to the slice to simulate memory allocation
		data = append(data, i)
	}
	// Prevent data from being GC'd too early
	fmt.Println("Length of data:", len(data))
}
