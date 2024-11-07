package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Set the number of goroutines to the number of CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Example data
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Sequential sum
	start := time.Now()
	sum := 0
	for _, value := range data {
		sum += value
	}
	fmt.Println("Sequential sum:", sum, "Time:", time.Since(start))

	// Parallel sum
	start = time.Now()
	var wg sync.WaitGroup
	partialSums := make([]int, runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := index; j < len(data); j += runtime.NumCPU() {
				partialSums[index] += data[j]
			}
		}(i)
	}

	wg.Wait()

	sum = 0
	for _, partialSum := range partialSums {
		sum += partialSum
	}
	fmt.Println("Parallel sum:", sum, "Time:", time.Since(start))
}
