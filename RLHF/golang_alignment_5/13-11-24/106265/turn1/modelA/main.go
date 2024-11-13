package main

import (
	"fmt"
	"time"
)

const (
	// Use a constant for the maximum value to be used in the arithmetic operations
	MAX_VALUE = 255 // uint8 maximum value
)

func main() {
	// Original arithmetic operations
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		result := originalAdd(i%MAX_VALUE, i%MAX_VALUE)
		//additional operation to force evaluation
		_ = result * result
	}
	end := time.Since(start)
	fmt.Println("Original Add:", end)

	// Bitwise optimized addition
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		result := optimizedAdd(uint8(i%MAX_VALUE), uint8(i%MAX_VALUE))
		_ = result * result
	}
	end = time.Since(start)
	fmt.Println("Optimized Add:", end)
}

// Original addition function
func originalAdd(a, b int) int {
	return a + b
}

// Optimized addition function using bit manipulation (for uint8 values)
func optimizedAdd(a, b uint8) uint8 {
	sum := a ^ b
	carry := (a & b) << 1
	// If carry is 0, we are done, otherwise, repeat the process with the carry
	for carry != 0 {
		temp := sum
		sum = carry ^ temp
		carry = (temp & carry) << 1
	}
	return sum
}
