package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	if n < 0 {
		return result
	}
	for i := 1; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
func main() {
	n := 1000
	fmt.Println("Factorial of", n, ":", factorial(n))
}
