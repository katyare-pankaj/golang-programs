package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Large amount of data to concatenate
	const dataSize = 1000000
	data := make([]string, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = fmt.Sprintf("data = %d ", i)
	}

	// Using bytes.Buffer for efficient concatenation
	var buffer bytes.Buffer
	for _, item := range data {
		buffer.WriteString(item)
	}

	// Convert the buffer to a string and print it
	result := buffer.String()
	fmt.Println(result)
}
