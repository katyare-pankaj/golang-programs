// package mathutils provides some basic mathematical utility functions.
package main

import "fmt"

// Add adds two numbers.
//
// It returns the sum of the two numbers.
//
// Example:
func main() {
	result := Add(10, 20)
	fmt.Println(result) // Output: 30
}

func Add(a, b float64) float64 {
	return a + b
}
