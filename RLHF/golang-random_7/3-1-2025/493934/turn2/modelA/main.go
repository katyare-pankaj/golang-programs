package main

import (
	"fmt"
	"strings"
	"testing"
)

// ReverseStringUsingLoop reverses a string using a for loop.
func ReverseStringUsingLoop(input string) string {
	var reversed string
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}

// ReverseStringUsingBuilder reverses a string using strings.Builder.
func ReverseStringUsingBuilder(input string) string {
	var builder strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		builder.WriteByte(input[i])
	}
	return builder.String()
}

func main() {
	// Example usage
	input := "Hello, World!"
	reversedLoop := ReverseStringUsingLoop(input)
	reversedBuilder := ReverseStringUsingBuilder(input)

	fmt.Println("Original:", input)
	fmt.Println("Reversed using loop:", reversedLoop)
	fmt.Println("Reversed using Builder:", reversedBuilder)

	// Run the benchmarks
	testing.Main()
}
