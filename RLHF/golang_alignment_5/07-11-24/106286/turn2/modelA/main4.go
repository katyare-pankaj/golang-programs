package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Set the maximum number of CPU cores to be used
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(data))

	var wg sync.WaitGroup
	wg.Add(len(data))

	// Parallel computation
	for i := range data {
		go func(index int) {
			result[index] = data[index] * data[index]
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Parallel Result:", result)
}
