package main

import "fmt"

// Define discount rates for different customer types
const (
	regularDiscountRateMondayToWednesday = 0.1
	regularDiscountRateThursdayToFriday  = 0.15
	regularDiscountRateSaturdaySunday    = 0.2
	premiumDiscountRate                  = 0.18
	vipDiscountRate                      = 0.25
)

func calculateDiscount(price float64, customerType string, dayOfWeek string) float64 {
	discountRate := 0.0

	switch customerType {
	case "regular":
		switch dayOfWeek {
		case "Monday", "Tuesday", "Wednesday":
			discountRate = regularDiscountRateMondayToWednesday
		case "Thursday", "Friday":
			discountRate = regularDiscountRateThursdayToFriday
		case "Saturday", "Sunday":
			discountRate = regularDiscountRateSaturdaySunday
		}
	case "premium":
		discountRate = premiumDiscountRate
	case "vip":
		discountRate = vipDiscountRate
	}

	return price * (1 - discountRate)
}

func main() {
	finalPrice := calculateDiscount(100.0, "regular", "Friday")
	fmt.Println("Final price after discount:", finalPrice)
}
