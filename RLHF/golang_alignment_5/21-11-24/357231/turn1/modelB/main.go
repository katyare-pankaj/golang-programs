package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Simulate some work
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	}

	// Start memory profiling
	mf, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer mf.Close()
	pprof.WriteHeapProfile(mf)
}
