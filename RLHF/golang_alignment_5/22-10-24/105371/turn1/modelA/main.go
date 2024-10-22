package main

import (
	"fmt"
)

// SparseMatrix represents a sparse matrix using a map for efficient storage
type SparseMatrix struct {
	rows int
	cols int
	data map[int]map[int]float64
}

// NewSparseMatrix creates a new sparse matrix
func NewSparseMatrix(rows, cols int) *SparseMatrix {
	return &SparseMatrix{
		rows: rows,
		cols: cols,
		data: make(map[int]map[int]float64),
	}
}

// Set sets a value at a specific row and column in the sparse matrix
func (m *SparseMatrix) Set(row, col int, value float64) {
	if m.data[row] == nil {
		m.data[row] = make(map[int]float64)
	}
	m.data[row][col] = value
}

// Get gets the value at a specific row and column in the sparse matrix
func (m *SparseMatrix) Get(row, col int) float64 {
	if rowData, ok := m.data[row]; ok {
		return rowData[col]
	}
	return 0
}

// Multiply performs matrix multiplication with another sparse matrix
func (m *SparseMatrix) Multiply(other *SparseMatrix) *SparseMatrix {
	if m.cols != other.rows {
		panic("Matrices cannot be multiplied due to incompatible dimensions.")
	}

	result := NewSparseMatrix(m.rows, other.cols)

	// Reuse row and col maps to avoid memory allocation
	var row map[int]float64
	var col map[int]float64

	for i := 0; i < m.rows; i++ {
		if rowData, ok := m.data[i]; ok {
			row = rowData
		} else {
			continue // Skip empty rows
		}
		for j := 0; j < other.cols; j++ {
			sum := 0.0
			for k := 0; k < m.cols; k++ {
				val1 := row[k]
				val2 := other.Get(k, j)
				sum += val1 * val2
			}
			if sum != 0 {
				if colData, ok := result.data[i]; !ok {
					col = make(map[int]float64)
					result.data[i] = col
				} else {
					col = colData
				}
				col[j] = sum
			}
		}
	}

	return result
}

func main() {
	// Create two sparse matrices
	m1 := NewSparseMatrix(3, 4)
	m1.Set(0, 1, 2.0)
	m1.Set(1, 0, 4.0)
	m1.Set(1, 2, 3.0)
	m1.Set(2, 3, 1.0)

	m2 := NewSparseMatrix(4, 2)
	m2.Set(0, 0, 5.0)
	m2.Set(0, 1, 6.0)
	m2.Set(1, 0, 7.0)
	m2.Set(3, 1, 8.0)

	// Multiply the matrices
	result := m1.Multiply(m2)

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
}
