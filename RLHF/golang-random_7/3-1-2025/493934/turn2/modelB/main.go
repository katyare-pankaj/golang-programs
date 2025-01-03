package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Example data
	data := []string{"Hello", ", ", "world", "!", " How", " are", " you", "?"}

	var buffer bytes.Buffer
	for _, part := range data {
		buffer.WriteString(part)
	}

	formattedString := buffer.String()
	fmt.Println(formattedString)
}
