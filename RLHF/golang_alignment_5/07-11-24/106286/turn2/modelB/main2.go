package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// Create a sparse matrix
	data := []float64{1.0, 2.0, 3.0, 4.0}
	rows := []int{0, 1, 0, 2}
	cols := []int{0, 1, 2, 1}
	S := mat64.NewSparse(3, 3, data, rows, cols)

	// Perform matrix multiplication with a dense matrix
	D := mat64.NewDense(3, 3, []float64{5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0})
	result := mat64.NewDense(3, 3, nil)
	S.Mul(result, D)

	fmt.Println("Sparse Matrix Result:")
	result.Print(10, 2)
}
