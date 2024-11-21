package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fibonacci calculates the nth Fibonacci number.
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	result := Fibonacci(n)
	fmt.Println("Fibonacci of", n, "is:", result)
}
