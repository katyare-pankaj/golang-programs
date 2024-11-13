package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func main() {
	// Your main app code here
}

func TestMemoryLeak(t *testing.T) {
	var leaks []*int
	for i := 0; i < 10000; i++ {
		data := new(int)
		*data = i
		leaks = append(leaks, data) // Accumulate leaks
	}
	// Defer a function to trigger garbage collection and check memory usage
	defer func() {
		runtime.GC()
		checkMemoryUsage(t)
	}()
	// Pause for a moment to allow garbage collection to happen
	time.Sleep(time.Second)
	// After the sleep, the memory usage should be high due to the memory leak
	checkMemoryUsage(t)
}

func checkMemoryUsage(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)

	// Add a threshold to fail the test if the memory usage exceeds a certain limit
	if m.Alloc > 100*1024*1024 { // 100MB threshold
		t.Fatalf("Memory usage exceeded limit: Alloc=%v MiB", bToMb(m.Alloc))
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
