package main

import (
	"fmt"
	"time"
)

const (
	MAX_VALUE = 15 // uint8 maximum value
)

func main() {
	// Original multiplication function
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		_ = originalMul(uint8(i%MAX_VALUE), uint8(i%MAX_VALUE))
	}
	end := time.Since(start)
	fmt.Println("Original Mul:", end)

	// Bitwise optimized multiplication
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		_ = optimizedMul(uint8(i%MAX_VALUE), uint8(i%MAX_VALUE))
	}
	end = time.Since(start)
	fmt.Println("Optimized Mul:", end)
}

// Original multiplication function
func originalMul(a, b uint8) uint8 {
	return a * b
}

// Bitwise optimized multiplication function (for uint8 values)
func optimizedMul(a, b uint8) uint8 {
	result := uint8(0)
	for b != 0 {
		if b&1 == 1 {
			result ^= a
		}
		a <<= 1 // Same as a *= 2
		b >>= 1 // Same as b /= 2
	}
	return result
}
