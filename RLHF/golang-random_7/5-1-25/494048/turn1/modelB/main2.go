package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	squares := []int{}
	for _, num := range numbers {
		squares = append(squares, num*num)
	}
	fmt.Println("Squares:", squares)
}
