package main

import (
	"math/rand"
	"sync"
	"testing"
)

// ParallelSearch performs a parallel search on a large genomic dataset
func ParallelSearch(data []string, target string, numGoroutines int) int {
	count := 0
	wg := &sync.WaitGroup{}

	chunkSize := len(data) / numGoroutines
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)

		go func(start, end int) {
			localCount := 0
			for j := start; j < end; j++ {
				if data[j] == target {
					localCount++
				}
			}
			count += localCount
			wg.Done()
		}(i, end)
	}

	wg.Wait()
	return count
}

func BenchmarkParallelSearch(b *testing.B) {
	// Generate a large genomic dataset for testing
	data := generateLargeGenomicData(1000000)
	target := "GCAT"
	numGoroutines := 4

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParallelSearch(data, target, numGoroutines)
	}
}

func TestParallelSearch(t *testing.T) {
	data := []string{"ATGC", "GCAT", "ATGC", "CGTA", "GCAT", "ATGC"}
	target := "GCAT"
	numGoroutines := 2

	expectedCount := 2
	result := ParallelSearch(data, target, numGoroutines)

	if result != expectedCount {
		t.Errorf("ParallelSearch failed: expected %d, got %d", expectedCount, result)
	}
}

func generateLargeGenomicData(size int) []string {
	// Generate a slice of random genomic strings for testing
	data := make([]string, size)
	for i := 0; i < size; i++ {
		data[i] = randomGenomicString(10)
	}
	return data
}

func randomGenomicString(n int) string {
	// Generate a random genomic string of length n
	var letters = []rune("ATGC")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
