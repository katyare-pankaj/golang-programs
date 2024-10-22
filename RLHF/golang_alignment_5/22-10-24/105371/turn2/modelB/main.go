package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CSRMatrix represents a sparse matrix in Compressed Sparse Row (CSR) format
type CSRMatrix struct {
	rows    int
	cols    int
	nnz     int // Number of non-zero elements
	data    []float64
	indices []int
	indptr  []int
}

// NewCSRMatrix creates a new CSR matrix
func NewCSRMatrix(rows, cols, nnz int) *CSRMatrix {
	return &CSRMatrix{
		rows:    rows,
		cols:    cols,
		nnz:     nnz,
		data:    make([]float64, nnz),
		indices: make([]int, nnz),
		indptr:  make([]int, rows+1),
	}
}

// Set sets a value at a specific row and column in the CSR matrix
func (m *CSRMatrix) Set(row, col int, value float64) {
	m.data[m.indptr[row]] = value
	m.indices[m.indptr[row]] = col
	m.indptr[row]++
}

// Multiply performs matrix multiplication with another CSR matrix using parallel processing
func (m *CSRMatrix) Multiply(other *CSRMatrix) *CSRMatrix {
	if m.cols != other.rows {
		panic("Matrices cannot be multiplied due to incompatible dimensions.")
	}

	result := NewCSRMatrix(m.rows, other.cols, m.rows*other.cols)

	// Create a wait group to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Launch a goroutine for each row of the result matrix
	for i := 0; i < m.rows; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < other.cols; j++ {
				sum := 0.0
				for k := m.indptr[row]; k < m.indptr[row+1]; k++ {
					val1 := m.data[k]
					val2 := other.Get(m.indices[k], j)
					sum += val1 * val2
				}
				if sum != 0 {
					result.Set(row, j, sum)
				}
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return result
}

// Get gets the value at a specific row and column in the CSR matrix
func (m *CSRMatrix) Get(row, col int) float64 {
	for i := m.indptr[row]; i < m.indptr[row+1]; i++ {
		if m.indices[i] == col {
			return m.data[i]
		}
	}
	return 0
}

func main() {
	// Set the seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// Generate a large, sparse dataset
	const rows = 10000
	const cols = 10000
	const nnz = 1000000 // Approximately 1% non-zero elements

	m1 := NewCSRMatrix(rows, cols, nnz)
	for i := 0; i < nnz; i++ {
		row := rand.Intn(rows)
		col := rand.Intn(cols)
		value := rand.Float64()
		m1.Set(row, col, value)
	}

	m2 := NewCSRMatrix(cols, rows, nnz)
	for i := 0; i < nnz; i++ {
		row := rand.Intn(cols)
		col := rand.Intn(rows)
		value := rand.Float64()
		m2.Set(row, col, value)
	}

	// Perform matrix multiplication
	start := time.Now()
	result := m1.Multiply(m2)
	duration := time.Since(start)

	// Print the result
	fmt.Println("Result:")
	for i := 0; i < result.rows; i++ {
		for j := 0; j < result.cols; j++ {
			val := result.Get(i, j)
			if val != 0 {
				fmt.Printf("m[%d][%d] = %.1f\n", i, j, val)
			}
		}
	}

	fmt.Printf("Matrix multiplication took %s\n", duration)
}
