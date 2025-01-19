package main

import (
	"fmt"
	"time"
	"unsafe"
)

// OptimizedData demonstrates cache-friendly struct alignment
type OptimizedData struct {
	A int32
	B int32
	C int32
	D int32 // Adjusted to align well in 64-byte cache line
}

func highFrequencyOperationWithStruct(data []OptimizedData) {
	for i := range data {
		data[i].A *= 2
		data[i].B *= 2
		data[i].C *= 2
		data[i].D *= 2
	}
}

func main() {
	start := time.Now()

	// Allocate a large slice of optimized data
	data := make([]OptimizedData, 1000000)

	fmt.Printf("Size of OptimizedData: %d bytes\n", unsafe.Sizeof(data[0]))

	// Run the high-frequency operation with optimized struct
	highFrequencyOperationWithStruct(data)

	fmt.Printf("Operation with optimized data took %s\n", time.Since(start))
}
