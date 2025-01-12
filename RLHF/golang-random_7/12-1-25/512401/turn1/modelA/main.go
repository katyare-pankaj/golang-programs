package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Generate some slices to profile
	slices := make([][]int, 1000)
	for i := range slices {
		slices[i] = make([]int, 10000)
	}

	// Start memory profiling
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

	// Optionally, wait for some time to allow garbage collection to happen
	time.Sleep(5 * time.Second)
}
