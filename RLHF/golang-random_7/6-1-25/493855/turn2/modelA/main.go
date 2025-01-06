package main

import (
	"fmt"
	"strings"
	"time"
)

const numConcatenations = 1000000
const stringLength = 20

func concatenateWithPlus() string {
	var result = ""
	for i := 0; i < stringLength; i++ {
		result += "h"
	}
	return result
}

func concatenateWithBuilder() string {
	var builder strings.Builder
	for i := 0; i < stringLength; i++ {
		builder.WriteString("h")
	}
	return builder.String()
}

func concatenateWithFmtSprintf() string {
	var result string
	for i := 0; i < stringLength; i++ {
		result = fmt.Sprintf("%s%s", result, "h")
	}
	return result
}

func main() {
	// Benchmark using the + operator
	var plusStart = time.Now()
	for i := 0; i < numConcatenations; i++ {
		_ = concatenateWithPlus()
	}
	var plusEnd = time.Now()
	fmt.Printf("Time taken using + operator: %v\n", plusEnd.Sub(plusStart))

	// Benchmark using strings.Builder
	var builderStart = time.Now()
	for i := 0; i < numConcatenations; i++ {
		_ = concatenateWithBuilder()
	}
	var builderEnd = time.Now()
	fmt.Printf("Time taken using strings.Builder: %v\n", builderEnd.Sub(builderStart))

	// Benchmark using fmt.Sprintf
	var fmtStart = time.Now()
	for i := 0; i < numConcatenations; i++ {
		_ = concatenateWithFmtSprintf()
	}
	var fmtEnd = time.Now()
	fmt.Printf("Time taken using fmt.Sprintf: %v\n", fmtEnd.Sub(fmtStart))
}
