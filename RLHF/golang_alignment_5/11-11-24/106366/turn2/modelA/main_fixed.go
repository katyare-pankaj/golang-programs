package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Low Cohesive Code
func calculateSumsLowCohesion(data [][]int) [][]int {
	results := make([][]int, len(data))
	wg := sync.WaitGroup{}

	for i, row := range data {
		wg.Add(1)
		go func(i int, row []int) {
			defer wg.Done()
			sums := make([]int, len(row))
			for j, num := range row {
				sums[j] = num + 10 // Some arbitrary operation
			}
			results[i] = sums
		}(i, row)
	}

	wg.Wait()
	return results
}

func calculateProductsLowCohesion(data [][]int) [][]int {
	results := make([][]int, len(data))
	wg := sync.WaitGroup{}

	for i, row := range data {
		wg.Add(1)
		go func(i int, row []int) {
			defer wg.Done()
			products := make([]int, len(row))
			for j, num := range row {
				products[j] = num * 2 // Some arbitrary operation
			}
			results[i] = products
		}(i, row)
	}
	wg.Wait()
	return results
}

// Highly Cohesive Code
func processDataHighCohesion(data [][]int, ops []func(int) int) [][]int {
	results := make([][]int, len(data))
	wg := sync.WaitGroup{}

	for i, row := range data {
		wg.Add(1)
		go func(i int, row []int) {
			defer wg.Done()
			processedRow := make([]int, len(row))
			for j, num := range row {
				for _, op := range ops {
					processedRow[j] = op(num)
				}
			}
			results[i] = processedRow
		}(i, row)
	}

	wg.Wait()
	return results
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dataSize := 10
	data := make([][]int, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			data[i][j] = j
		}
	}

	// Low Cohesion Example
	start := time.Now()
	sums := calculateSumsLowCohesion(data)
	products := calculateProductsLowCohesion(data)
	fmt.Println("Low Cohesion Result: Sum =", sums, ", Product =", products)
	fmt.Println("Low Cohesion Time:", time.Since(start))

	// High Cohesion Example
	start = time.Now()
	ops := []func(int) int{func(x int) int { return x + 10 }, func(x int) int { return x * 2 }}
	processedData := processDataHighCohesion(data, ops)
	fmt.Println("High Cohesion Results:", processedData)
	fmt.Println("High Cohesion Time:", time.Since(start))
}
