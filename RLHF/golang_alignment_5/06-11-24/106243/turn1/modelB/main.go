package main

import (
	"os"
	"runtime/pprof"
)

func busyLoop(n int) {
	for i := 0; i < n; i++ {
		// Simulate some work
	}
}

func main() {
	// Start CPU profiling
	f, err := os.Create("cpuprofile")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Call the function that you want to profile
	busyLoop(100000000)

	// Print the profiling results
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}
