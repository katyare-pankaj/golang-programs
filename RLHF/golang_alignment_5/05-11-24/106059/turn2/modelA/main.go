package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/05-11-24/106001/turn2/modelA/calculator"
)

func main() {
	calc := calculator.Calculator{}
	result := calc.Add(20.0, 30.0)
	fmt.Println("Result from main:", result)
}
