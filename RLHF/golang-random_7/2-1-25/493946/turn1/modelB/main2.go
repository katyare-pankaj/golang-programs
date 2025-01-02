package main

import (
	"runtime"
	"sync"
)

func analyzeData(data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Process the data
}

func main() {
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5}
	numGoroutines := runtime.NumCPU()
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go analyzeData(data, &wg)
	}

	wg.Wait()
}
