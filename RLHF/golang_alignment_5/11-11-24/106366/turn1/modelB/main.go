// Highly Cohesive Matrix Multiplication
package main

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])

	if cols1 != rows2 {
		panic("Number of columns in first matrix must equal number of rows in second matrix.")
	}

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}

	// Calculate the result matrix using nested loops
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

// Low Cohesion Matrix Multiplication
func multiplyMatricesLowCohesion(matrix1 [][]int, matrix2 [][]int) [][]int {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	rows2 := len(matrix2)
	cols2 := len(matrix2[0])

	if cols1 != rows2 {
		panic("Number of columns in first matrix must equal number of rows in second matrix.")
	}

	// Perform various steps to calculate the result matrix
	calculateResultMatrix(matrix1, matrix2, rows1, cols1, rows2, cols2)

	return getResultMatrix(rows1, cols2)
}

func calculateResultMatrix(matrix1 [][]int, matrix2 [][]int, rows1 int, cols1 int, rows2 int, cols2 int) {
	// Calculate the result matrix using nested loops
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				updateResultMatrix(i, j, k, matrix1, matrix2)
			}
		}
	}
}

func updateResultMatrix(i int, j int, k int, matrix1 [][]int, matrix2 [][]int) {
	result[i][j] += matrix1[i][k] * matrix2[k][j]
}

func getResultMatrix(rows1 int, cols2 int) [][]int {
	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}
	return result
}
