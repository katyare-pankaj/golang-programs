package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func createBigObjects() {
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 10000) // Allocate a large slice of bytes
		// Perform some operations with the large array
		time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	}
}

func main() {
	// Generate memory profile
	f, err := os.Create("memory.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)

	// Simulate some work and create big objects
	createBigObjects()
	fmt.Println("Memory profiling complete. Open memory.prof with go tool pprof to analyze.")
}
