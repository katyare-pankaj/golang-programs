package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Low Cohesive Code
func calculateSumLowCohesion(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func calculateProductLowCohesion(numbers []int) int {
	product := 1
	for _, num := range numbers {
		product *= num
	}
	return product
}

// Highly Cohesive Code
func processNumbersHighCohesion(numbers []int, ops []func(int) int) []int {
	results := make([]int, len(ops))
	wg := sync.WaitGroup{}
	for i, op := range ops {
		wg.Add(1)
		go func(i int, op func(int) int) {
			defer wg.Done()
			result := 0
			for _, num := range numbers {
				result = op(num)
			}
			results[i] = result
		}(i, op)
	}
	wg.Wait()
	return results
}
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available cores    numElements := 1_000_000
	numbers := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		numbers[i] = i + 1
	}
	// Low Cohesion Example
	start := time.Now()
	sum := calculateSumLowCohesion(numbers)
	product := calculateProductLowCohesion(numbers)
	fmt.Println("Low Cohesion Result: Sum =", sum, ", Product =", product)
	fmt.Println("Low Cohesion Time:", time.Since(start))
	// High Cohesion Example
	start = time.Now()
	ops := []func(int) int{calculateSum, calculateProduct}
	results := processNumbersHighCohesion(numbers, ops)
	fmt.Println("High Cohesion Results:", results)
	fmt.Println("High Cohesion Time:", time.Since(start))
}
