package main

import (
	"fmt"
)

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}
	outer(fruits)
}

func outer(fruits []string) {
	for i, fruit := range fruits {
		fmt.Println("Outer loop: index =", i, ", fruit =", fruit)
		inner(fruits, i, fruit)
	}
}

func inner(fruits []string, i int, fruit string) {
	fmt.Println("Inner loop: index =", i, ", fruit =", fruit)
	// If you modify i or fruit here, it will not affect the outer scope.
}
