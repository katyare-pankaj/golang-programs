package main

import "fmt"

func calculateDiscount(price float64, customerType string, dayOfWeek string) float64 {
	discountRate := 0.0

	switch customerType {
	case "regular":
		switch dayOfWeek {
		case "Monday", "Tuesday", "Wednesday":
			discountRate = 0.1
		case "Thursday", "Friday":
			discountRate = 0.15
		case "Saturday", "Sunday":
			discountRate = 0.2
		}
	case "premium":
		discountRate = 0.18
	case "vip":
		discountRate = 0.25
	}

	return price * (1 - discountRate)
}

func main() {
	finalPrice := calculateDiscount(100.0, "regular", "Friday")
	fmt.Println("Final price after discount:", finalPrice)
}
