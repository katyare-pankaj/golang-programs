package main

import (
	"fmt"
	"time"
)

// Order represents a shipping order
type Order struct {
	ID       int
	Weight   float64
	Distance float64
}

// ShippingCostCalculator calculates shipping cost for an order
func ShippingCostCalculator(order Order) float64 {
	// Simple cost calculation logic for demonstration
	baseCost := 2.5
	weightCost := order.Weight * 0.75
	distanceCost := order.Distance * 0.1
	return baseCost + weightCost + distanceCost
}

func processOrdersBatch(orders []Order) {
	for _, order := range orders {
		cost := ShippingCostCalculator(order)
		fmt.Printf("Order ID: %d, Shipping Cost: $%.2f\n", order.ID, cost)
	}
}

func main() {
	// Sample orders data
	orders := []Order{
		{ID: 1, Weight: 2.0, Distance: 100.0},
		{ID: 2, Weight: 5.0, Distance: 50.0},
		// Add more order data as needed
	}

	// Schedule the batch process to run every 24 hours
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Running scheduled batch shipping cost calculation...")
			processOrdersBatch(orders)
			fmt.Println("Batch processing completed.")
		}
	}
}
