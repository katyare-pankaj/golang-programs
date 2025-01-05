package main

import (
	"fmt"
)

func formatDataWithStringManipulation(data []int) string {
	var result string = "["
	for i, value := range data {
		result += fmt.Sprint(value)
		if i < len(data)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	formattedData := formatDataWithStringManipulation(data)
	fmt.Println(formattedData) // Output: [1, 2, 3, 4, 5]
}
