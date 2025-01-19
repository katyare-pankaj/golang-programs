package main

import (
	"fmt"
	"sync"
)

// OptimizeData struct with padding to prevent false sharing
type OptimizeData struct {
	value int64
	_     [7]int64 // Padding to occupy a full cache line (assuming 64-byte cache lines)
}

func main() {
	const workers = 4
	var wg sync.WaitGroup
	data := make([]OptimizeData, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				// Perform some computation and update in a thread-safe manner
				data[idx].value += int64(j)
			}
		}(i)
	}

	wg.Wait()

	// Compute total to ensure the compiler doesn't optimize away the loop
	var total int64
	for i := 0; i < workers; i++ {
		total += data[i].value
	}
	fmt.Println("Total:", total)
}
