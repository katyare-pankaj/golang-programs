package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix is a type alias for a 2D slice.
type Matrix [][]float64

// transposeMatrix transposes the given matrix in-place.
func transposeMatrix(mat Matrix) {
	rows := len(mat)
	if rows == 0 {
		return
	}
	cols := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
}

// newMatrix creates a zero-initialized matrix of the given size.
func newMatrix(rows, cols int) Matrix {
	mat := make(Matrix, rows)
	for i := range mat {
		mat[i] = make([]float64, cols)
	}
	return mat
}

// randomMatrix generates a matrix filled with random float64 numbers.
func randomMatrix(rows, cols int) Matrix {
	mat := newMatrix(rows, cols)
	for i := range mat {
		for j := range mat[i] {
			mat[i][j] = rand.Float64()
		}
	}
	return mat
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Define matrix dimensions
	rows, cols := 1000, 1000

	// Create and populate the matrix with random data
	matrix := randomMatrix(rows, cols)

	// Measure execution time of transposition
	start := time.Now()
	transposeMatrix(matrix)
	elapsed := time.Since(start)

	// Output the time taken for transposition
	fmt.Printf("Matrix transposition complete in %s\n", elapsed)
}
