package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Example data
	data := []string{"Hello", ", ", "world", "!", " How", " are", " you", "?"}

	// Preallocate capacity for bytes.Buffer
	buffer := bytes.NewBuffer(make([]byte, 0, 100)) // Initial capacity of 100 bytes

	// Concatenate strings using bytes.Buffer
	for _, part := range data {
		buffer.WriteString(part)
	}

	formattedString1 := buffer.String()
	fmt.Println("Concatenated using bytes.Buffer:", formattedString1)

	// Reset the buffer for reuse
	buffer.Reset()

	// Preallocate capacity for strings.Builder
	builder := strings.Builder{}
	builder.Grow(100) // Initial capacity of 100 bytes

	// Concatenate strings using strings.Builder
	for _, part := range data {
		builder.WriteString(part)
	}

	formattedString2 := builder.String()
	fmt.Println("Concatenated using strings.Builder:", formattedString2)
}
