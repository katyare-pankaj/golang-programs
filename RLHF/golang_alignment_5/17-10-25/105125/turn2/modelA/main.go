package main

import (
	"fmt"
)

// Order represents a shipping order
type Order struct {
	ID           int
	Weight       float64
	ShippingCost float64
}

func aggregateAndCalculateAverageCost(orders []Order) map[string]float64 {
	weightCategories := map[string][]float64{
		"Light":  []float64{},
		"Medium": []float64{},
		"Heavy":  []float64{},
	}

	// Aggregate data by weight category
	for _, order := range orders {
		var category string
		if order.Weight <= 2 {
			category = "Light"
		} else if order.Weight <= 10 {
			category = "Medium"
		} else {
			category = "Heavy"
		}
		weightCategories[category] = append(weightCategories[category], order.ShippingCost)
	}

	// Calculate average cost for each category
	averages := make(map[string]float64)
	for category, costs := range weightCategories {
		total := 0.0
		for _, cost := range costs {
			total += cost
		}
		if len(costs) > 0 {
			averages[category] = total / float64(len(costs))
		} else {
			averages[category] = 0
		}
	}
	return averages
}

func main() {
	// Sample orders data
	orders := []Order{
		{ID: 1, Weight: 1.5, ShippingCost: 10.50},
		{ID: 2, Weight: 8.0, ShippingCost: 18.00},
		{ID: 3, Weight: 4.0, ShippingCost: 12.00},
		{ID: 4, Weight: 12.0, ShippingCost: 25.00},
		{ID: 5, Weight: 6.0, ShippingCost: 15.00},
		// Add more order data
	}

	fmt.Println("Running batch data aggregation for average shipping costs...")
	averageCosts := aggregateAndCalculateAverageCost(orders)

	for category, averageCost := range averageCosts {
		fmt.Printf("Average Shipping Cost for %s category: $%.2f\n", category, averageCost)
	}
}
