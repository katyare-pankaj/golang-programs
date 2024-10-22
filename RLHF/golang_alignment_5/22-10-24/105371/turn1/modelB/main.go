package main

import (
	"fmt"
)

// SparseMatrix using a hash map
type SparseMatrix struct {
	data map[int]int
	rows int
	cols int
}

// NewSparseMatrix creates a new SparseMatrix
func NewSparseMatrix(rows, cols int) *SparseMatrix {
	return &SparseMatrix{
		data: make(map[int]int),
		rows: rows,
		cols: cols,
	}
}

// SetElement sets an element in the sparse matrix
func (m *SparseMatrix) SetElement(row, col, val int) {
	key := row*m.cols + col
	m.data[key] = val
}

// GetElement retrieves an element from the sparse matrix
func (m *SparseMatrix) GetElement(row, col int) int {
	key := row*m.cols + col
	return m.data[key]
}

// IsSet checks if an element is set in the sparse matrix
func (m *SparseMatrix) IsSet(row, col int) bool {
	key := row*m.cols + col
	_, ok := m.data[key]
	return ok
}

// SparseBitmap using an array of uint8
type SparseBitmap struct {
	data []uint8
}

// NewSparseBitmap creates a new SparseBitmap with the specified number of bits
func NewSparseBitmap(bits int) *SparseBitmap {
	return &SparseBitmap{
		data: make([]uint8, (bits+7)/8),
	}
}

// SetBit sets the specified bit in the bitmap
func (b *SparseBitmap) SetBit(bit int) {
	byteIndex := bit / 8
	bitIndex := uint(bit % 8)
	b.data[byteIndex] |= 1 << bitIndex
}

// ClearBit clears the specified bit in the bitmap
func (b *SparseBitmap) ClearBit(bit int) {
	byteIndex := bit / 8
	bitIndex := uint(bit % 8)
	b.data[byteIndex] &= ^(1 << bitIndex)
}

// TestBit tests the specified bit in the bitmap and returns true if it is set
func (b *SparseBitmap) TestBit(bit int) bool {
	byteIndex := bit / 8
	bitIndex := uint(bit % 8)
	return (b.data[byteIndex] & (1 << bitIndex)) != 0
}
func main() {
	// Using SparseMatrix
	fmt.Println("Sparse Matrix Example:")
	sparseMatrix := NewSparseMatrix(10, 10)
	sparseMatrix.SetElement(2, 3, 42)
	sparseMatrix.SetElement(7, 5, 13)

	if val := sparseMatrix.GetElement(2, 3); val != 42 {
		fmt.Println("Error: Expected 42, got", val)
	}

	if val := sparseMatrix.GetElement(7, 5); val != 13 {
		fmt.Println("Error: Expected 13, got", val)
	}

	fmt.Println("Element (2,3) is set:", sparseMatrix.IsSet(2, 3))
	fmt.Println("Element (8,8) is set:", sparseMatrix.IsSet(8, 8)) // false
	// Using SparseBitmap
	fmt.Println("\nSparse Bitmap Example:")
	const numBits = 1000
	sparseBitmap := NewSparseBitmap(numBits)
	sparseBitmap.SetBit(3)
	sparseBitmap.SetBit(17)
	sparseBitmap.SetBit(42)

	if !sparseBitmap.TestBit(3) {
		fmt.Println("Error: Bit 3 should be set")
	}

	if sparseBitmap.TestBit(numBits - 1) {
		fmt.Println("Error: Bit", numBits-1, "should be clear")
	}
	fmt.Println("Bit 17 is set:", sparseBitmap.TestBit(17))
	fmt.Println("Bit 42 is set:", sparseBitmap.TestBit(42))
}
