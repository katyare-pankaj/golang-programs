package main

import (
	"fmt"
)

func generateStrings(pattern string, start int, end int) []string {
	strings := make([]string, 0)

	for i := start; i <= end; i++ {
		str := fmt.Sprintf(pattern, i)
		strings = append(strings, str)
	}
	return strings
}

func main() {
	pattern := "String%03d"
	start := 1
	end := 10
	result := generateStrings(pattern, start, end)
	fmt.Println(result)
}
