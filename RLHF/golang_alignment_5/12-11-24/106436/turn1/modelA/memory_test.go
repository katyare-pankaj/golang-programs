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

func TestMemoryUsage(t *testing.T) {
	for i := 0; i < 10; i++ {
		func() {
			// Simulate a memory-intensive operation here
			data := make([]byte, 1024*1024*10) // 10MB allocation
			defer func() {
				runtime.GC()
				checkMemoryUsage(t)
			}()
			// Perform your social app logic here
			time.Sleep(time.Second)
		}()
	}
}

func checkMemoryUsage(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)

	// Add your thresholds here to fail the test if memory usage exceeds limits
	if m.Alloc > 100*1024*1024 { // 100MB threshold
		t.Fatalf("Memory usage exceeded limit: Alloc=%v MiB", bToMb(m.Alloc))
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
