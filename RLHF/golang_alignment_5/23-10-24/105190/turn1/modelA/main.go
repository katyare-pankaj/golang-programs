package main

import (
	"fmt"
	"sync"
)

// CalculateStatistic calculates a statistic from genomic data in parallel
func CalculateStatistic(data []int) int {
	sum := 0
	wg := &sync.WaitGroup{}

	// Partition the data into smaller chunks for parallel processing
	chunkSize := len(data) / 4 // Assuming 4 cores
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)

		go func(start, end int) {
			partialSum := 0
			for j := start; j < end; j++ {
				partialSum += data[j]
			}
			sum += partialSum
			wg.Done()
		}(i, end)
	}

	wg.Wait()
	return sum
}

func main() {
	// Example usage:
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := CalculateStatistic(data)
	fmt.Println("Calculated Statistic:", result)
}
