package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
)

const (
	sliceSize     = 10000000               // large slice size
	numGoroutines = 4                      // number of goroutines to use
	sleepDuration = 100 * time.Millisecond // time for goroutines to work
)

func main() {
	// allocate large slice to profile memory usage
	largeSlice := make([]int, sliceSize)

	// Create a progress bar
	bar := pb.New(len(largeSlice)).SetRefreshRate(time.Second)
	bar.Start()

	// Update bar as go routine works on the slice
	go func() {
		for i := 0; i < len(largeSlice); i++ {
			time.Sleep(sleepDuration) // simulation of some work
			largeSlice[i] = rand.Intn(1000)
			bar.Increment()
		}
		bar.Finish()
	}()

	// Measure memory usage every second using a goroutine
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Alloc = %6vMB HeapSys = %6vMB HeapAlloc = %6vMB GC HeapAllocs = %d  HeapFrees = %d\n",
				m.Alloc/1024/1024, m.HeapSys/1024/1024, m.HeapAlloc/1024/1024, m.GCCount, m.HeapFrees)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Process the large slice using multiple goroutines
	chunkSize := sliceSize / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		go processSlice(largeSlice[start:end], &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	time.Sleep(2 * time.Second) // Wait a bit to observe the final memory usage
}

func processSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range slice {
		time.Sleep(sleepDuration) // simulate some work
		num++                     // some processing on the slice
	}
}
