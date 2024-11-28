package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/28-11-24/389103/turn1/modelA/calculator"
)

// To test the program
func main() {

	cal := calculator.NewCalculator()

	cal.Add(10)
	cal.Subtract(10)
	cal.Multiply(10)
	cal.Divide(10)
	cal.GetMemory()
	cal.clearMemory()

	fmt.Println("The code is successfully compiled.")

}
