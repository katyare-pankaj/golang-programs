package main

import (
	"fmt"
	"sync"
	"time"
)

func processNumbersWithGoroutines(start, end int, wg *sync.WaitGroup, results []int, index int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		results[i-start+index*10000000] = i * i // Simulate CPU-bound work
	}
}

func main() {
	const goroutineCount = 8
	start := 1
	end := 100000000
	results := make([]int, end-start)

	var wg sync.WaitGroup
	step := (end - start) / goroutineCount

	startTime := time.Now()
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go processNumbersWithGoroutines(start+i*step, start+(i+1)*step, &wg, results, i)
	}

	wg.Wait()
	elapsed := time.Since(startTime)

	fmt.Printf("Goroutine-based calculation took %s\n", elapsed)
	fmt.Println("Sample output:", results[:5])
}
