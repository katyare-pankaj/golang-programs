package ui

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/05-11-24/106001/turn2/modelB/arithmetic"
)

func Calculate() {
	var op string
	var a, b float64

	fmt.Print("Enter operation (+, -, etc.): ")
	fmt.Scan(&op)

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	var result float64
	var err error
	switch op {
	case "+":
		result, err = arithmetic.Add(a, b)
	case "-":
		result, err = arithmetic.Subtract(a, b)
	// Add more operations
	default:
		fmt.Println("Invalid operation.")
		return
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}
}
