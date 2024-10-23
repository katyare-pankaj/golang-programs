package main

import (
	"fmt"
	"sync"
	"testing"
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

// TestCalculateStatistic ensures the correctness of CalculateStatistic function
func TestCalculateStatistic(t *testing.T) {
	testCases := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Small dataset",
			data:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "Large dataset",
			data:     make([]int, 1000000),
			expected: 499999500000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateStatistic(tc.data)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkCalculateStatistic(b *testing.B) {
	data := make([]int, 10000000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateStatistic(data)
	}
}

func main() {
	// Example usage:
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := CalculateStatistic(data)
	fmt.Println("Calculated Statistic:", result)
}
