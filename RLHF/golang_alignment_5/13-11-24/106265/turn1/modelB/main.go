package main

import (
	"fmt"
)

func bitwiseSqrt(x uint32) uint32 {
	if x < 2 {
		return x
	}

	y := x / 2
	for y*y > x {
		y = (y + x/y) / 2
	}

	return y
}

func main() {
	num := uint32(16)
	result := bitwiseSqrt(num)
	fmt.Println("Square root of", num, "is:", result) // Output: Square root of 16 is: 4
}
