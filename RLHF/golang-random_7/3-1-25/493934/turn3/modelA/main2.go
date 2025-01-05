package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example data: a large slice of strings
	data := []string{"Hello", ", ", "world", "!", " How", " are", " you", "?"}

	// Calculate the total length of the strings
	totalLength := 0
	for _, part := range data {
		totalLength += len(part)
	}

	// Create a strings.Builder and preallocate memory
	var builder strings.Builder
	builder.Grow(totalLength)

	// Concatenate strings into the builder
	for _, part := range data {
		builder.WriteString(part)
	}

	formattedString := builder.String()
	fmt.Println(formattedString)
}
