package main

import (
	"fmt"
	"sync"
)

// Multiplies two matrices
func multiplyMatrices(A [][]int, B [][]int) [][]int {
	rowsA := len(A)
	colsA := len(A[0])
	rowsB := len(B)
	colsB := len(B[0])

	// Ensure matrices can be multiplied
	if colsA != rowsB {
		panic("Matrices cannot be multiplied")
	}

	// Resultant matrix
	C := make([][]int, rowsA, rowsA)
	for i := range C {
		C[i] = make([]int, colsB, colsB)
	}

	// Parallely calculate each element of matrix C
	var wg sync.WaitGroup
	for i := 0; i < rowsA; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < colsB; j++ {
				for k := 0; k < colsA; k++ {
					C[row][j] += A[row][k] * B[k][j]
				}
			}
		}(i)
	}

	wg.Wait()
	return C
}

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	B := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	C := multiplyMatrices(A, B)

	fmt.Println("Resulting matrix:")
	for _, row := range C {
		fmt.Println(row)
	}
}
