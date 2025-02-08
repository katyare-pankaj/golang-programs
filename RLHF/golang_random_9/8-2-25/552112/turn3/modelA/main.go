package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix is a type alias for a 2D slice.
type Matrix [][]float64

// newMatrix initializes a zero-valued matrix with specified rows and columns.
func newMatrix(rows, cols int) Matrix {
	mat := make(Matrix, rows)
	row := make([]float64, rows*cols)
	for i := range mat {
		mat[i], row = row[:cols], row[cols:] // efficient contiguous allocation
	}
	return mat
}

// transposeMatrix transposes the input matrix 'm' to the output matrix 'res'.
// Assumes res is already initialized with dimensions cols(m)xrows(m).
func transposeMatrix(m, res Matrix) {
	rows := len(m)
	cols := len(m[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[j][i] = m[i][j]
		}
	}
}

// randomMatrix generates a matrix filled with random float64 values.
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
	m := randomMatrix(rows, cols)
	transposed := newMatrix(cols, rows)

	// Measure execution time of the transpose operation
	start := time.Now()
	transposeMatrix(m, transposed)
	elapsed := time.Since(start)

	// Output the time taken for transposition
	fmt.Printf("Matrix transposition complete in %s\n", elapsed)

	// Uncomment below to print part of the transposed matrix
	// fmt.Println(transposed[:3][:3]) // prints part of the transposed matrix, useful for debug
}
