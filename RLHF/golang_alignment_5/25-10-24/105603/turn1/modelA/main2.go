package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Set the number of goroutines to use
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(data))

	var wg sync.WaitGroup
	wg.Add(len(data))

	for i := 0; i < len(data); i++ {
		go func(index int) {
			result[index] = data[index] * data[index]
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Print the results
	fmt.Println(result)
}
