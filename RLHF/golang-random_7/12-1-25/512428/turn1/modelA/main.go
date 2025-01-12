package main

import (
	"fmt"
)

func generateSequence(min, max int, format string) []string {
	var sequence []string
	for i := min; i <= max; i++ {
		sequence = append(sequence, fmt.Sprintf(format, i))
	}
	return sequence
}

func main() {
	min := 1
	max := 5
	format := "%04d" // Format the number as a four-digit string
	sequence := generateSequence(min, max, format)

	for _, str := range sequence {
		fmt.Println(str)
	}
}
