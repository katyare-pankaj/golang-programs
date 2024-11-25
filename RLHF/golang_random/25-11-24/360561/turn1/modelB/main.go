package main

import (
	"fmt"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

// PCA performs Principal Component Analysis on a given dataset.
func PCA(data [][]float64, numComponents int) ([][]float64, []float64) {
	// Check if the number of components is less than or equal to the number of features
	if numComponents <= 0 || numComponents > len(data[0]) {
		return nil, nil
	}
	// Create a covariance matrix
	covMatrix := mat64.NewDense(len(data[0]), len(data[0]), nil)
	for _, row := range data {
		covMatrix.MulTransposed(covMatrix, mat64.NewDense(1, len(row), row))
	}
	covMatrix.Scale(1.0/float64(len(data)), covMatrix)
	// Eigenvalue decomposition of the covariance matrix
	eigenvalues, eigenvectors := covMatrix.Eigen(nil)
	// Sort eigenvalues and eigenvectors in descending order
	floats.SortDescend(eigenvalues.RawVector().Data)
	eigenvectors.Sort(func(i, j int) bool {
		return eigenvalues.At(i, 0) > eigenvalues.At(j, 0)
	})
	// Extract the top numComponents eigenvectors
	principalComponents := mat64.NewDense(numComponents, len(data[0]), nil)
	for i := 0; i < numComponents; i++ {
		principalComponents.SetRow(i, eigenvectors.RawRowView(i))
	}
	// Calculate the transformed data
	transformedData := mat64.NewDense(len(data), numComponents, nil)
	for i, row := range data {
		transformedData.Mul(transformedData, principalComponents)
	}
	return transformedData.RawMatrix().Data, eigenvalues.RawVector().Data
}
func main() {
	// Sample dataset for demonstration
	data := [][]float64{
		{2.5, 0.5},
		{1.4, 1.5},
		{3.1, 1.0},
		{2.0, 2.0},
		{1.0, 1.0},
	}
	// Perform PCA with 1 component
	transformedData, eigenvalues := PCA(data, 1)
	fmt.Println("Transformed Data:")
	fmt.Println(transformedData)
	fmt.Println("\nEigenvalues:")
	fmt.Println(eigenvalues)
}
