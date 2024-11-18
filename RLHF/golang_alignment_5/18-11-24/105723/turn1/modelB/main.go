package main

import (
	"fmt"
)

const maximumUint64 = ^uint64(0)

func factorial(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialIterative(n uint64) uint64 {
	if n > maximumUint64/2 {
		return 0 // Reset factorial of numbers too big for uint64, they would overflow
	}

	if n == 0 || n == 1 {
		return 1
	}
	var result uint64 = 1
	for i := uint64(2); i <= n; i++ {
		result *= i
		if result > maximumUint64/i {
			return 0 // Overflow would occur if we did i*result in the next step
		}
	}
	return result
}

func main() {
	const n uint64 = 17
	fmt.Println("Factorial of", n, "using recursion:", factorial(n))
	fmt.Println("Factorial of", n, "using iteration:", factorialIterative(n))

	// If you uncomment these, you will run into raised stack again, underscores demonstrate this is too big
	// for a uint64 to store
	// hiBigNumber := uint64(10000)
	// fmt.Println("Factorial of", hiBigNumber, "using recursion:", factorial(hiBigNumber))
	// fmt.Println("Factorial of", hiBigNumber, "using iteration:", factorialIterative(hiBigNumber))
}
