package main

import (
	"fmt"
	"time"
)

func processNumbersSingleThreaded(start, end int) []int {
	results := make([]int, end-start)
	for i := start; i < end; i++ {
		results[i-start] = i * i // Simulate CPU-bound work
	}
	return results
}

func main() {
	start := 1
	end := 100000000

	startTime := time.Now()
	results := processNumbersSingleThreaded(start, end)
	elapsed := time.Since(startTime)

	fmt.Printf("Single-threaded calculation took %s\n", elapsed)
	fmt.Println("Sample output:", results[:5])
}
