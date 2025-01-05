package main

import (
	"fmt"
)

func main() {

	value := 123.45678
	formatted := fmt.Sprintf("Value: %.2f", value) // Output: Value: 123.46
	fmt.Println(formatted)

	var result string
	for i := 0; i < 100; i++ {
		result += fmt.Sprint(i) + " "
	}
	fmt.Println(result) // Output: 0 1 2 3 ... 99
}
