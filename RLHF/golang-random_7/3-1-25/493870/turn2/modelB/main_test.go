package main

import (
	"fmt"
	"strings"
	"testing"
)

func reverseStringForLoop(input string) string {
	result := ""
	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}
	return result
}

func reverseStringBuilder(input string) string {
	var builder strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		builder.WriteByte(input[i])
	}
	return builder.String()
}

func BenchmarkReverseStringForLoop(b *testing.B) {
	input := "Hello, World!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseStringForLoop(input)
	}
}

func BenchmarkReverseStringBuilder(b *testing.B) {
	input := "Hello, World!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseStringBuilder(input)
	}
}

func main() {
	input := "Hello, World!"
	fmt.Println("Reversed string using for loop:", reverseStringForLoop(input))
	fmt.Println("Reversed string using strings.Builder:", reverseStringBuilder(input))
}
