package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Example data: a large slice of strings
	data := []string{"Hello", ", ", "world", "!", " How", " are", " you", "?"}

	// Calculate the total length of the strings
	totalLength := 0
	for _, part := range data {
		totalLength += len(part)
	}

	// Preallocate the buffer to minimize memory reallocations
	var buffer bytes.Buffer
	buffer.Grow(totalLength)

	// Concatenate strings into the buffer
	for _, part := range data {
		buffer.WriteString(part)
	}

	formattedString := buffer.String()
	fmt.Println(formattedString)
}
