package main

import (
	"fmt"
	"math"
	"time"

	"github.com/sugarme/gotch"
)

// NaiveMatrixMultiply performs a naive matrix multiplication
func NaiveMatrixMultiply(A, B [][]float32) [][]float32 {
	rowsA := len(A)
	colsA := len(A[0])
	rowsB := len(B)
	colsB := len(B[0])

	if colsA != rowsB {
		panic("Number of columns in A must equal number of rows in B.")
	}

	C := make([][]float32, rowsA)
	for i := 0; i < rowsA; i++ {
		C[i] = make([]float32, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

// EfficientMatrixMultiply uses NumGO for efficient matrix multiplication
func EfficientMatrixMultiply(A, B *gotch.Tensor) *gotch.Tensor {
	return A.Matmul(B)
}

func main() {
	// Initialize Gotch
	gotch.Init()
	defer gotch.Close()

	// Generate large random matrices for demonstration
	size := 1000
	A := gotch.Randn(size, size, gotch.Float)
	B := gotch.Randn(size, size, gotch.Float)

	// Naive Matrix Multiplication
	start := time.Now()
	naiveResult := NaiveMatrixMultiply(A.MustData().([]float32), B.MustData().([]float32))
	naiveDuration := time.Since(start)

	// Efficient Matrix Multiplication using NumGO
	start = time.Now()
	efficientResult := EfficientMatrixMultiply(A, B)
	efficientDuration := time.Since(start)

	// Verify the results are equal
	gotch.AssertEqual(naiveResult, efficientResult.MustData().([]float32))

	fmt.Printf("Naive Matrix Multiplication Time: %s\n", naiveDuration)
	fmt.Printf("Efficient Matrix Multiplication Time using NumGO: %s\n", efficientDuration)
	fmt.Printf("Speedup: %.2fx\n", math.Ceil(float64(naiveDuration)/float64(efficientDuration)))
}
