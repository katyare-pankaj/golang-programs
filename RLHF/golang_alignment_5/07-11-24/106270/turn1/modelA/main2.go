// Without Abstraction (Optimized for Merging)

package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func main() {

	resultAdd := Add(5, 3)
	fmt.Println("5 + 3 =", resultAdd)

	resultMultiply := Multiply(5, 3)
	fmt.Println("5 * 3 =", resultMultiply)
}
