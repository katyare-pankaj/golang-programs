package main

import (
	"fmt"
	"strings"
	"time"
)

const numIterations = 100000
const inputString = "hello world"

func toUpperWithStrings() string {
	return strings.ToUpper(inputString)
}

func toUpperWithLoop() string {
	var result string
	for _, char := range inputString {
		result += string(strings.ToUpperRune(char))
	}
	return result
}

func toUpperWithBuilder() string {
	var builder strings.Builder
	for _, char := range inputString {
		builder.WriteRune(strings.ToUpperRune(char))
	}
	return builder.String()
}

func main() {
	// Benchmark using strings.ToUpper
	var toUpperStart = time.Now()
	for i := 0; i < numIterations; i++ {
		_ = toUpperWithStrings()
	}
	var toUpperEnd = time.Now()
	fmt.Printf("Time taken using strings.ToUpper: %v\n", toUpperEnd.Sub(toUpperStart))

	// Benchmark using loop to convert characters
	var loopStart = time.Now()
	for i := 0; i < numIterations; i++ {
		_ = toUpperWithLoop()
	}
	var loopEnd = time.Now()
	fmt.Printf("Time taken using loop: %v\n", loopEnd.Sub(loopStart))

	// Benchmark using strings.Builder
	var builderStart = time.Now()
	for i := 0; i < numIterations; i++ {
		_ = toUpperWithBuilder()
	}
	var builderEnd = time.Now()
	fmt.Printf("Time taken using strings.Builder: %v\n", builderEnd.Sub(builderStart))
}
