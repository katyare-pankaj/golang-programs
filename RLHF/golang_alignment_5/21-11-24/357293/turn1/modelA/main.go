package main

import (
	"fmt"
)

// FactorizeNumber decomposes the number into prime factors.
func factorizeNumber(n int) []int {
	var factors []int
	x := uint64(n)

	for i := uint64(2); i*i <= x; i++ {
		for x%i == 0 {
			factors = append(factors, int(i))
			x /= i
		}
	}
	if x > 1 {
		factors = append(factors, int(x))
	}

	return factors
}

// Main function to demonstrate factorizeNumber.
func main() {
	number := 100
	fmt.Println("Factors of", number, ":", factorizeNumber(number))
}
