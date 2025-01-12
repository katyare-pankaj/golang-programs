package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Create a large slice
	slices := make([][]int, 1000)
	for i := range slices {
		slices[i] = make([]int, 10000)
		for j := range slices[i] {
			slices[i][j] = i + j
		}
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

	// Set slices to nil to free memory
	slices = nil
	time.Sleep(5 * time.Second)

	// Create another slice to show the difference in memory usage
	newSlices := make([][]int, 1000)
	for i := range newSlices {
		newSlices[i] = make([]int, 10000)
		for j := range newSlices[i] {
			newSlices[i][j] = i + j
		}
	}

	// Generate a second memory profile
	f2, err := os.Create("memprofile2.prof")
	if err != nil {
		fmt.Println("could not create memory profile:", err)
		return
	}
	defer f2.Close()
	if err := pprof.WriteHeapProfile(f2); err != nil {
		fmt.Println("could not write memory profile:", err)
	}
	fmt.Println("memory profile written to memprofile2.prof")

	// Free memory by setting newSlices to nil
	newSlices = nil
	time.Sleep(5 * time.Second)
}
