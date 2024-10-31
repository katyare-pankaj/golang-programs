// Factorial calculates the factorial of a non-negative number.

package main

import "fmt"

func main() {
	res := Factorial(5)
	fmt.Println("Factorial :", res)
}
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
