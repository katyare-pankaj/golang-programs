package main

import (
	"fmt"
	"sync"
	"time"
)

// Order struct represents an order in the system
type Order struct {
	OrderID  string
	Customer string
}

// ProcessOrder performs the order processing
func ProcessOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done() //notify the WaitGroup that this goroutine is done
	fmt.Println("Processing order:", order.OrderID, "for customer:", order.Customer)
	// Simulate processing time (2 seconds)
	time.Sleep(2 * time.Second)
	fmt.Println("Order:", order.OrderID, "processed successfully.")
}

// GenerateOrders generates orders with a given delay between orders
func GenerateOrders(wg *sync.WaitGroup, orders chan Order) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		order := Order{
			OrderID:  fmt.Sprintf("ORDER_%03d", i),
			Customer: fmt.Sprintf("Customer_%03d", i),
		}
		orders <- order
		time.Sleep(1 * time.Second) // Delay between order generations
	}
	close(orders) // Close the orders channel to signal no more orders
}

func main() {
	orders := make(chan Order)
	wg := &sync.WaitGroup{}

	wg.Add(1) // for GenerateOrders
	go GenerateOrders(wg, orders)
	// We can set the number of workers (goroutines) based on your system's capabilities
	// For simplicity, let's use 3 workers in this example.
	numWorkers := 3

	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // for each worker
		go func() {
			for order := range orders {
				ProcessOrder(order, wg)
			}
			wg.Done()
		}()
	}

	fmt.Println("Order processing started...")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Order processing completed.")
}
