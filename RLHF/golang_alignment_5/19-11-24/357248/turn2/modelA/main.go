package main

import (
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)

func factorialIterative(n int) *big.Int {
	if n < 0 {
		return zero
	}
	result := big.NewInt(1)
	if n == 0 || n == 1 {
		return result
	}

	for i := 1; i <= n; i++ {
		temp := big.NewInt(0)
		temp.Mul(result, big.NewInt(int64(i)))
		result = temp
	}
	return result
}

func factorialRecursive(n int) *big.Int {
	if n < 0 {
		return zero
	}

	return factorialHelper(n, big.NewInt(1))
}

func factorialHelper(n int, result *big.Int) *big.Int {
	if n == 0 || n == 1 {
		return result
	}

	temp := big.NewInt(0)
	temp.Mul(result, big.NewInt(int64(n)))
	return factorialHelper(n-1, temp)
}

func main() {
	n := 20
	fmt.Println("Factorial of", n, "using iteration:", factorialIterative(n))
	fmt.Println("Factorial of", n, "using recursion:", factorialRecursive(n))

	n = 100
	fmt.Println("Factorial of", n, "using iteration:", factorialIterative(n))
	fmt.Println("Factorial of", n, "using recursion:", factorialRecursive(n))
}
