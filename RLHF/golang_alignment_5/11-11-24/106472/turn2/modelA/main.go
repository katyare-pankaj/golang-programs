package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// Generate a large dataset for processing
	data := generateLargeDataset(10000000)

	// Perform processing and calculate statistics
	start := time.Now()
	min, max, sum := processData(data)
	duration := time.Since(start)

	fmt.Println("Processing time:", duration)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Sum:", sum)
	fmt.Println("Memory Usage:", runtime.MemStats{}.Alloc)
}

func generateLargeDataset(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000)
	}
	return data
}

func processData(data []int) (min, max, sum int) {
	// Using slices for in-place processing to reduce memory allocation
	min = data[0]
	max = data[0]
	sum = data[0]

	for _, value := range data[1:] {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
		sum += value
	}
	return
}
