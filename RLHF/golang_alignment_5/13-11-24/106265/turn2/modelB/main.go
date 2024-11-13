package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = originalMul(uint8(i%16), uint8(i%16))
	}
	end := time.Since(start)
	fmt.Println()
	fmt.Println("Original Mul:", end)

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = optimizedMul(uint8(i%16), uint8(i%16))
	}
	end = time.Since(start)
	fmt.Println("Optimized Mul:", end)
}

// Original multiplication function
func originalMul(a, b uint8) uint8 {
	return a * b
}

// Optimized multiplication function using bit manipulation (for uint8 values)
func optimizedMul(a, b uint8) uint8 {
	result := uint8(0)
	for b != 0 {
		if b&1 == 1 {
			result += a
		}
		a <<= 1 // Multiply a by 2 using left shift
		b >>= 1 // Divide b by 2 using right shift
	}
	return result
}
