package main

import (
	"fmt"
	"sync"
)

func processSlice(data []int, wg *sync.WaitGroup) {
	defer wg.Done() // This is executed when the goroutine completes

	for _, value := range data {
		// For demonstration purposes, print each value
		// In real-world applications, replace this with actual processing logic
		fmt.Println("Processing:", value)
	}
}

func main() {
	var wg sync.WaitGroup

	// Example slices of data to be processed
	dataSlices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, data := range dataSlices {
		wg.Add(1) // Increment the WaitGroup counter
		go processSlice(data, &wg)
	}

	wg.Wait() // Block until all goroutines have finished
	fmt.Println("All slices processed.")
}
