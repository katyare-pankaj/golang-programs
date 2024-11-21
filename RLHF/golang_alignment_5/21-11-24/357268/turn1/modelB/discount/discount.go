// discount.go
package discount

import (
	"go-programs/RLHF/golang_alignment_5/21-11-24/357268/turn1/modelB/calculator"
)

func CalculateDiscount(price, discountRate float64) float64 {
	discount := calculator.Multiply(price, discountRate)
	return calculator.Subtract(price, discount)
}
