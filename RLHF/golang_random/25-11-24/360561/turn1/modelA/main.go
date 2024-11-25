package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

// PCA performs PCA on a dataset to reduce its dimensionality while preserving as much variance as possible.
func PCA(data *mat.Dense, numComponents int) (*mat.Dense, error) {
	// Ensure data has at least as many rows as columns
	if data.Rows < data.Cols {
		return nil, fmt.Errorf("data must have at least as many rows as columns")
	}

	// Compute the covariance matrix
	cov := mat.NewDense(data.Cols, data.Cols)
	if err := mat.Cov(cov, data); err != nil {
		return nil, err
	}

	// Compute the SVD of the covariance matrix
	svd := mat.NewSVDRandomized(data.Cols, data.Cols, numComponents)
	if err := svd.Do(cov); err != nil {
		return nil, err
	}

	// The columns of svd.V are the principal components
	components := mat.NewDense(data.Cols, numComponents)
	for i := 0; i < numComponents; i++ {
		components.SetCol(i, svd.V.Col(i))
	}

	// Project the data onto the principal components
	projectedData := mat.NewDense(data.Rows, numComponents)
	if err := mat.Gemv(mat.NoTrans, 1, data, components, 0, projectedData); err != nil {
		return nil, err
	}

	return projectedData, nil
}

func main() {
	// Example usage
	data := mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})
	numComponents := 1

	projectedData, err := PCA(data, numComponents)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Projected data:\n%v", mat.Formatted(projectedData))
}
