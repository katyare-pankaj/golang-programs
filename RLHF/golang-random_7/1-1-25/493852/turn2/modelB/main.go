package main

import (
	"fmt"
)

func stringConcatenation(a, b, c string) string {
	return a + b + c
}

func main() {
	str1 := "Hello, "
	str2 := "world!"
	str3 := " How are you?"
	result := stringConcatenation(str1, str2, str3)
	fmt.Println(result)
}
