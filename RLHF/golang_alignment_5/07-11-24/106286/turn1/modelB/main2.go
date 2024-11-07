// Example of parallelized matrix multiplication
package main

import (
	"runtime"
)

func main() {
	// Set the number of CPUs to use
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create two matrices
	rowsA, colsA := 3, 4
	rowsB, colsB := 4, 2
	a := make([][]float64, rowsA)
	for i := range a {
		a[i] = make([]float64, colsA)
	}
	b := make([][]float64, rowsB)
	for i := range b {
		b[i] = make([]float64, colsB)
	}
	// Result matrix
	c := make([][]float64, rowsA)
	for i := range c {
		c[i] = make([]float64, colsB)
	}

	// Initialize the matrices
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			a[i][j] = float64(i*colsA + j)
		}
	}
	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			b[i][j] = float64(i*colsB + j)
		