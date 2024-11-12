package main

import (
	"fmt"
)

const (
	dataSize     = 4 // Size of an int in bytes (on 32-bit or 64-bit systems)
	unrollFactor = 4 // Unrolling factor to reduce loop overhead
)

// Serializes a slice of integers into a byte slice using loop unrolling
func serializeIntegers(data []int, buf []byte) {
	for i := 0; i < len(data); i += unrollFactor {
		for j := 0; j < unrollFactor; j++ {
			idx := i + j
			if idx >= len(data) {
				break
			}
			// Unrolling the loop: serialize 4 integers at once
			buf[idx*dataSize] = byte(data[idx] >> 24)
			buf[idx*dataSize+1] = byte(data[idx] >> 16)
			buf[idx*dataSize+2] = byte(data[idx] >> 8)
			buf[idx*dataSize+3] = byte(data[idx])
		}
	}
}

// Deserializes a byte slice back into a slice of integers using loop unrolling
func deserializeIntegers(buf []byte, data []int) {
	for i := 0; i < len(data); i += unrollFactor {
		for j := 0; j < unrollFactor; j++ {
			idx := i + j
			if idx >= len(data) {
				break
			}
			// Unrolling the loop: deserialize 4 integers at once
			data[idx] = int(buf[idx*dataSize])<<24 |
				int(buf[idx*dataSize+1])<<16 |
				int(buf[idx*dataSize+2])<<8 |
				int(buf[idx*dataSize+3])
		}
	}
}

func main() {
	// Create a slice of integers for serialization
	const numDataElements = 1000
	data := make([]int, numDataElements)
	buf := make([]byte, numDataElements*dataSize)

	// Initialize data for serialization
	for i := range data {
		data[i] = i * 2 // Just an example of initializing data
	}

	// Serialize the data
	serializeIntegers(data, buf)

	// Deserialize the data back into the slice
	deserializeIntegers(buf, data)

	// Verify deserialization
	for i := range data {
		if data[i] != i*2 {
			fmt.Println("Deserialization failed!")
			return
		}
	}
	fmt.Println("Serialization/Deserialization successful!")
}
