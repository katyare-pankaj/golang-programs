package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Order represents an order in the e-commerce platform
type Order struct {
	ID          int
	PlacedAt    time.Time
	ProcessedAt time.Time
}

// OrderProcessor handles the processing of orders
type OrderProcessor struct {
	orders []Order
}

// NewOrderProcessor creates a new OrderProcessor
func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{}
}

// AddOrder adds an order to the processor's queue
func (op *OrderProcessor) AddOrder(order Order) {
	op.orders = append(op.orders, order)
}

// ProcessOrders processes all the orders in the queue
func (op *OrderProcessor) ProcessOrders() {
	for _, order := range op.orders {
		// Simulate processing time
		processingTime := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(processingTime)
		order.ProcessedAt = time.Now()
	}
}

// AnalyzeProcessingTime analyzes the time taken to process orders
func AnalyzeProcessingTime(orders []Order) {
	totalProcessingTime := time.Duration(0)
	for _, order := range orders {
		processingTime := order.ProcessedAt.Sub(order.PlacedAt)
		totalProcessingTime += processingTime
	}

	averageProcessingTime := totalProcessingTime / time.Duration(len(orders))
	fmt.Printf("Average processing time: %s\n", averageProcessingTime)
}

func main() {
	// Create a new OrderProcessor
	orderProcessor := NewOrderProcessor()

	// Generate sample orders for the sprint
	for i := 1; i <= 20; i++ {
		order := Order{
			ID:       i,
			PlacedAt: time.Now().Add(-time.Duration(rand.Intn(10)+1) * time.Minute),
		}
		orderProcessor.AddOrder(order)
	}

	// Process the orders in the sprint
	orderProcessor.ProcessOrders()

	// Analyze the processing time for the sprint
	AnalyzeProcessingTime(orderProcessor.orders)
}
