package main

import (
	"fmt"
	"sync"
	"time"
)

// Low Cohesion Matrix Multiplication
func multiplyMatricesLowCohesion(matrix1 [][]int, matrix2 [][]int) [][]int {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])

	if cols1 != rows2 {
		panic("Number of columns in first matrix must equal number of rows in second matrix.")
	}

	// Create the result matrix
	result := getResultMatrix(rows1, cols2)

	// Perform various steps to calculate the result matrix
	calculateResultMatrix(matrix1, matrix2, rows1, cols1, rows2, cols2, result)

	return result
}

func calculateResultMatrix(matrix1 [][]int, matrix2 [][]int, rows1 int, cols1 int, rows2 int, cols2 int, result [][]int) {
	// Calculate the result matrix using nested loops
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				updateResultMatrix(i, j, k, matrix1, matrix2, result)
			}
		}
	}
}

func updateResultMatrix(i int, j int, k int, matrix1 [][]int, matrix2 [][]int, result [][]int) {
	result[i][j] += matrix1[i][k] * matrix2[k][j]
}

func getResultMatrix(rows1 int, cols2 int) [][]int {
	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}
	return result
}

// High Cohesion - Separate function for each independent task
func multiplyMatricesHighCohesion(matrix1, matrix2 [][]int) [][]int {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	cols2 := len(matrix2[0])

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}

	var wg sync.WaitGroup
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				for k := 0; k < cols1; k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}(i, j)
		}
	}
	wg.Wait()

	return result
}

func main() {
	// Define two sample matrices
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrix2 := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// Measure time for highly cohesive matrix multiplication
	start := time.Now()
	resultHighCohesion := multiplyMatricesHighCohesion(matrix1, matrix2)
	fmt.Println("High Cohesion Result:", resultHighCohesion)
	fmt.Println("High Cohesion Time:", time.Since(start))

	// Measure time for low cohesion matrix multiplication
	start = time.Now()
	resultLowCohesion := multiplyMatricesLowCohesion(matrix1, matrix2)
	fmt.Println("Low Cohesion Result:", resultLowCohesion)
	fmt.Println("Low Cohesion Time:", time.Since(start))
}
