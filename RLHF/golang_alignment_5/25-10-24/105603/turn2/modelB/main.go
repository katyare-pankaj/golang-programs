package main

import (
	"fmt"
	"time"

	"github.com/sugarme/gotch"
)

func main() {
	// Initialize Gotch
	gotch.Init()
	defer gotch.Close()

	// Set the dimensions of the matrices
	rows1 := 1000
	cols1 := 2000
	rows2 := 2000
	cols2 := 3000

	// Generate random input matrices
	a := gotch.Randn(rows1, cols1, gotch.Float32)
	b := gotch.Randn(rows2, cols2, gotch.Float32)

	// Original implementation using native Go slices and loops
	start := time.Now()
	c1 := classicMatrixMultiplication(a, b)
	duration1 := time.Since(start)

	// Optimized implementation using NumGO
	start = time.Now()
	c2 := optimizedMatrixMultiplication(a, b)
	duration2 := time.Since(start)

	// Print the results and execution time
	fmt.Println("Classic Matrix Multiplication:")
	fmt.Println(c1)
	fmt.Println("Time:", duration1)

	fmt.Println("\nOptimized Matrix Multiplication using NumGO:")
	fmt.Println(c2)
	fmt.Println("Time:", duration2)
}

// Classic matrix multiplication using native Go slices and loops
func classicMatrixMultiplication(a, b gotch.Tensor) gotch.Tensor {
	// Assert that matrix multiplication is possible
	if a.Size(1) != b.Size(0) {
		panic("Matrices cannot be multiplied due to incorrect dimensions!")
	}

	resultRows := a.Size(0)
	resultCols := b.Size(1)
	c := gotch.NewEmptyTensor([]int{resultRows, resultCols}, gotch.Float32)

	// Perform matrix multiplication using nested loops
	for i := 0; i < resultRows; i++ {
		for j := 0; j < resultCols; j++ {
			for k := 0; k < a.Size(1); k++ {
				c.Set2(i, j, c.Get2(i, j)+a.Get2(i, k)*b.Get2(k, j))
			}
		}
	}

	return c
}

// Optimized matrix multiplication using NumGO
func optimizedMatrixMultiplication(a, b gotch.Tensor) gotch.Tensor {
	// Perform matrix multiplication using NumGO's optimized functions
	c := a.MM(b)
	return c
}
