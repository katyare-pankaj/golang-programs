package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ParallelMap[T any, U any](input []T, fn func(T) U, numWorkers int) []U {
	output := make([]U, len(input))
	chunks := make([][]T, numWorkers)
	chunkSize := (len(input) + numWorkers - 1) / numWorkers
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		chunks[i/chunkSize] = input[i:end]
	}
	resultCh := make(chan map[int]U, numWorkers)
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(chunk []T, workerId int) {
			defer wg.Done()
			results := make(map[int]U)
			for j, v := range chunk {
				results[j+workerId*chunkSize] = fn(v)
			}
			resultCh <- results
		}(chunks[i], i)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	for resultMap := range resultCh {
		for k, v := range resultMap {
			output[k] = v
		}
	}
	return output
}

func main() {
	numWorkers := runtime.NumCPU()
	fmt.Println("Using", numWorkers, "workers")

	// Example usage with numbers.
	nums := make([]int, 10000000)
	for i := range nums {
		nums[i] = i + 1
	}

	squared := ParallelMap(nums, func(n int) int { return n * n }, numWorkers)
	fmt.Println("First 10 squared values:", squared[:10])

	sum := Reduce(squared, 0, func(a, b int) int { return a + b })
	fmt.Println("Sum of all squared values:", sum)
}
