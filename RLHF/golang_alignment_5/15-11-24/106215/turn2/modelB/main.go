package main

import (
	"fmt"
	"runtime"
	"time"
)

func quicksort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	runtime.GOMAXPROCS(1) // Set GOMAXPROCS to 1 for single-threaded execution
	const size = 100000
	arr := make([]int, size)

	// Generate random array
	for i := range arr {
		arr[i] = size - i
	}

	start := time.Now()
	quicksort(arr, 0, size-1)
	elapsed := time.Since(start)

	fmt.Printf("Sorted array: %v\n", arr)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
