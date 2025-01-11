package main

import (
	"fmt"
)

// Custom iterator interface
type customIterator struct {
	current int
	limit   int
}

// Next implements the iterator interface
func (c *customIterator) Next() bool {
	if c.current < c.limit {
		c.current++
		return true
	}
	return false
}

// Value implements the iterator interface
func (c *customIterator) Value() int {
	return c.current - 1
}

func main() {
	// Example 1: Iterating over a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Iterating over a slice:")
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example 2: Iterating over a map
	people := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 22}
	fmt.Println("\nIterating over a map:")
	for key, age := range people {
		fmt.Printf("%s is %d years old.\n", key, age)
	}

	// Example 3: Iterating over an array
	fruits := [5]string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Println("\nIterating over an array:")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Example 4: Iterating over a string
	sentence := "Hello, Go!"
	fmt.Println("\nIterating over a string:")
	for _, char := range sentence {
		fmt.Printf("%c", char)
	}
	fmt.Println()

	// Example 5: Custom iterator
	limit := 10
	fmt.Println("\nIterating over a custom iterator:")
	iter := customIterator{current: 1, limit: limit}
	for iter.Next() {
		fmt.Println(iter.Value())
	}
}
