package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SparseData represents a sparse dataset using a map for efficient storage
type SparseData struct {
	data map[int]map[int]float64
}

// NewSparseData creates a new sparse data structure
func NewSparseData() *SparseData {
	return &SparseData{
		data: make(map[int]map[int]float64),
	}
}

// Set sets a value at a specific row and column in the sparse data
func (sd *SparseData) Set(row, col int, value float64) {
	if sd.data[row] == nil {
		sd.data[row] = make(map[int]float64)
	}
	sd.data[row][col] = value
}

// Get gets the value at a specific row and column in the sparse data
func (sd *SparseData) Get(row, col int) float64 {
	if rowData, ok := sd.data[row]; ok {
		return rowData[col]
	}
	return 0
}

// ComputeMean computes the mean of the non-zero values in the sparse data
func (sd *SparseData) ComputeMean() float64 {
	sum := 0.0
	count := 0

	// Iterate through the sparse data to calculate sum and count
	for _, rowData := range sd.data {
		for _, value := range rowData {
			sum += value
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func main() {
	// Generate a large sparse dataset for demonstration purposes
	const rows = 100000
	const cols = 10000
	const sparsity = 0.01 // Percentage of non-zero elements

	rand.Seed(time.Now().UnixNano())

	sparseData := NewSparseData()

	// Populate the sparse data with random values
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if rand.Float64() < sparsity {
				value := rand.Float64() * 100
				sparseData.Set(row, col, value)
			}
		}
	}

	// Measure the time taken to compute the mean
	startTime := time.Now()
	mean := sparseData.ComputeMean()
	duration := time.Since(startTime)

	fmt.Println("Sparse Data Mean:", mean)
	fmt.Println("Time taken to compute mean:", duration)
}
