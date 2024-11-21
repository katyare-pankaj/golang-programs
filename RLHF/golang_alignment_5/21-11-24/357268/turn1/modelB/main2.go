// main.go
package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/21-11-24/357268/turn1/modelB/discount"
)

func main() {
	originalPrice := 100.0
	discountRate := 0.15
	finalPrice := discount.CalculateDiscount(originalPrice, discountRate)

	fmt.Println("Original Price:", originalPrice)
	fmt.Println("Discount Rate:", discountRate)
	fmt.Println("Final Price:", finalPrice)
}
