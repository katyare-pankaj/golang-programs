package main

import (
	"fmt"
	"strings"
	"time"
)

const numIterations = 1000000
const stringLength = 20

func convertToUpperManual(input string) string {
	result := ""
	for _, char := range input {
		result += string(char + 'A' - 'a')
	}
	return result
}

func main() {
	testString := strings.Repeat("a", stringLength)

	// Benchmark strings.ToUpper
	start := time.Now()
	for i := 0; i < numIterations; i++ {
		_ = strings.ToUpper(testString)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken using strings.ToUpper: %v\n", elapsed)

	// Benchmark converting to uppercase manually
	start = time.Now()
	for i := 0; i < numIterations; i++ {
		_ = convertToUpperManual(testString)
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken using manual conversion: %v\n", elapsed)

	// Benchmark using strings.Builder to build the uppercase string
	start = time.Now()
	for i := 0; i < numIterations; i++ {
		var builder strings.Builder
		builder.Grow(stringLength)
		for _, char := range testString {
			builder.WriteRune(char + 'A' - 'a')
		}
		_ = builder.String()
	}
	elapsed = time.Since(start)
	fmt.Printf("Time taken using strings.Builder: %v\n", elapsed)
}
