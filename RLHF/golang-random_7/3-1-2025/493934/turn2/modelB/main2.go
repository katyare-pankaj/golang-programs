package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example data
	data := []string{"Hello", ", ", "world", "!", " How", " are", " you", "?"}

	var builder strings.Builder
	for _, part := range data {
		builder.WriteString(part)
	}

	formattedString := builder.String()
	fmt.Println(formattedString)
}
