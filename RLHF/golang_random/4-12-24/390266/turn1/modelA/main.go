package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/4-12-24/390251/turn1/modelA/utils"
)

func main() {

	sum := utils.CalculateSum(10, 20)

	fmt.Println("Sum :", sum)

	sum2 := utils.calculateSum(10, 20)

	fmt.Println("Sum :", sum2)

}
