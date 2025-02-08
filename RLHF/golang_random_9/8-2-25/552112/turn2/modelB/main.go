package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func elementWiseMultiply(a, b [][]float64) [][]float64 {
	result := make([][]float64, len(a))
	for i := range a {
		result[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			result[i][j] = a[i][j] * b[i][j]
		}
	}
	return result
}
func createMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}
func main() {
	rand.Seed(time.Now().UnixNano())
	rows, cols := 1000, 1000 // Change these values to test performance with larger matrices
	mat1 := createMatrix(rows, cols)
	mat2 := createMatrix(rows, cols)
	start := time.Now()
	result := elementWiseMultiply(mat1, mat2)
	elapsed := time.Since(start)
	fmt.Printf("Element-wise multiplication complete in: %s\n", elapsed)
	fmt.Println("Memory usage:", runtime.MemStats{}.Alloc/1024/1024, "MB")
}
