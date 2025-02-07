package main

import (
	"fmt"
	"sync"
	"time"
)

// DeliveryOrder represents a delivery order with an ID and a delivery distance.
type DeliveryOrder struct {
	ID       int
	Distance int // in kilometers
}

// ProcessOrder simulates processing a delivery order.
func ProcessOrder(order DeliveryOrder, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// Simulate processing time with sleep
	time.Sleep(time.Millisecond * 100)

	// Example processing: Just printing the order details
	fmt.Printf("Processed order ID %d: Delivery distance %d km\n", order.ID, order.Distance)
}

func main() {
	// Example list of delivery orders
	deliveryOrders := []DeliveryOrder{
		{ID: 1, Distance: 50},
		{ID: 2, Distance: 75},
		{ID: 3, Distance: 120},
		{ID: 4, Distance: 30},
		{ID: 5, Distance: 90},
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Add the number of orders to process to the WaitGroup's counter
	wg.Add(len(deliveryOrders))

	// Launch a goroutine to process each delivery order
	for _, order := range deliveryOrders {
		// Pass the WaitGroup by reference
		go ProcessOrder(order, &wg)
	}

	// Wait for all goroutines to finish processing
	wg.Wait()
	fmt.Println("All delivery orders processed.")
}
