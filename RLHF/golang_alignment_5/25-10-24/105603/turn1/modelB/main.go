package main

import (
	"math/rand"
	"sync"
)

// Generate a random vector of size 'n' using SIMD optimizations
func generateRandomVector(n int) []float64 {
	vec := make([]float64, n)
	for i := 0; i < n; i++ {
		vec[i] = rand.Float64()
	}
	return vec
}

// Perform matrix multiplication concurrently using multiple goroutines
func matrixMultiplyConcurrently(a, b [][]float64) [][]float64 {
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]float64, colsB)
	}

	var wg sync.WaitGroup
	wg.Add(rowsA * colsB)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			go func(i, j int) {
				for k := 0; k < colsA; k++ {
					result[i][j] += a[i][k] * b[k][j]
				}
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait()
	return result
}
