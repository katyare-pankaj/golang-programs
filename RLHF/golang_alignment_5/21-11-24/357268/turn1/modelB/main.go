package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/21-11-24/357268/turn1/modelB/calculator"
)

func main() {
	sum := calculator.Add(10, 20)
	difference := calculator.Subtract(30, 15)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
