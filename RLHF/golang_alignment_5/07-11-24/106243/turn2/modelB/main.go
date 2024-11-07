package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	// Enable garbage collection debugging information
	runtime.SetGCPercent(100)

	// Simulate some work with memory allocation
	doSomeWork()

	// Generate the memory profile
	runtime.GC() // get up-to-date statistics
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("could not create memory profile:", err)
		return
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile:", err)
		return
	}

	fmt.Println("Memory profile generated.")
}

func doSomeWork() {
	const numIterations = 1000000

	// Leaky loop: Allocate memory without freeing it
	for i := 0; i < numIterations; i++ {
		newSlice := make([]byte, 1024)
		// Do something with the slice
		_ = newSlice
	}
}
