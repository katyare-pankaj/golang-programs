package main

import (
	"fmt"
	"sync"
)

func parallelMap(input []int, mapFunc func(int) int) []int {
	output := make([]int, len(input))
	var wg sync.WaitGroup

	for i, v := range input {
		wg.Add(1)
		go func(i, v int) {
			defer wg.Done()
			output[i] = mapFunc(v)
		}(i, v)
	}

	wg.Wait()
	return output
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	squared := parallelMap(numbers, func(x int) int { return x * x })
	fmt.Println(squared) // Output: [1 4 9 16 25]
}
