package main

import (
	"fmt"
	"strconv"
)

func convertStringsToInts(input []string) []int {
	var result []int

	for _, s := range input {
		if num, err := strconv.Atoi(s); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	strSlice := []string{"1", "2", "three", "4", "5", "a", "7", "b"}
	intSlice := convertStringsToInts(strSlice)
	fmt.Println(intSlice) // Output will be [1 2 4 5 7]
}
