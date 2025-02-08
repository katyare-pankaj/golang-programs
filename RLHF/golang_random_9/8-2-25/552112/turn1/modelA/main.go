package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix is a type alias for a 2D slice.
type Matrix [][]float64

// newMatrix creates a zero-initialized matrix of the given size.
func newMatrix(rows, cols int) Matrix {
	mat := make(Matrix, rows)
	for i := range mat {
		// Preallocate slices to reduce allocations within the inner loop
		mat[i] = make([]float64, cols)
	}
	return mat
}

// multiplyMatrices performs matrix multiplication for a * b and stores the result in res.
// Note: The number of columns in 'a' should equal the number of rows in 'b'.
func multiplyMatrices(a, b, res Matrix) {
	// Reducing allocations by storing initial lengths
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			// Directly assign multiplication iterating on k index for each matrix element.
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += a[i][k] * b[k][j]
			}
			res[i][j] = sum
		}
	}
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
	// Example of using the matrix multiplication function
	rand.Seed(time.Now().UnixNano())

	m1 := randomMatrix(500, 300) // 500 rows, 300 cols
	m2 := randomMatrix(300, 400) // 300 rows, 400 cols
	result := newMatrix(500, 400)

	start := time.Now()
	multiplyMatrices(m1, m2, result)
	elapsed := time.Since(start)

	fmt.Printf("Matrix multiplication complete in %s\n", elapsed)
}
