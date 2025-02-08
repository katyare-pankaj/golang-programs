package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix is a type alias for a 2D slice of float64.
type Matrix [][]float64

// newMatrix initializes a zero-valued matrix with specified rows and columns.
func newMatrix(rows, cols int) Matrix {
	mat := make(Matrix, rows)
	row := make([]float64, rows*cols)
	for i := range mat {
		mat[i], row = row[:cols], row[cols:] // efficient data allocation to reduce GC pressure
	}
	return mat
}

// elementWiseMultiply performs element-wise multiplication between two matrices a and b.
// It stores the result in the res matrix.
// Assumes a, b, and res have the same dimensions.
func elementWiseMultiply(a, b, res Matrix) {
	for i := range a {
		for j := range a[i] {
			res[i][j] = a[i][j] * b[i][j]
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

	// Create and populate the matrices with random data
	m1 := randomMatrix(rows, cols)
	m2 := randomMatrix(rows, cols)
	product := newMatrix(rows, cols)

	// Measure execution time of element-wise multiplication
	start := time.Now()
	elementWiseMultiply(m1, m2, product)
	elapsed := time.Since(start)

	// Output the time taken for multiplication
	fmt.Printf("Element-wise multiplication complete in %s\n", elapsed)

	// Uncomment below to print part of the matrix
	// fmt.Println(product[:3][:3]) // prints part of the matrix, useful for debug
}
