package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Enable maximum number of cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Sample data
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(data))

	// Parallel computation using goroutines and channels
	var wg sync.WaitGroup
	wg.Add(len(data))
	for i := range data {
		go func(index int) {
			result[index] = data[index] * data[index]
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("Parallel Result:", result)
}
