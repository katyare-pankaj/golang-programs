package main

import (
	"fmt"
)

func formatDataWithSprintf(data []int) string {
	return fmt.Sprintf("[%v]", data)
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	formattedData := formatDataWithSprintf(data)
	fmt.Println(formattedData) // Output: [1 2 3 4 5]
}
