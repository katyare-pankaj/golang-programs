package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// Order represents a supply chain order
type Order struct {
	ID             int
	ProcessingTime int // Time taken to process the order
	ArrivalTime    int // Time the order arrives
}

// OrderQueue implements a priority queue for Orders based on ArrivalTime
type OrderQueue []*Order

func (pq OrderQueue) Len() int { return len(pq) }

func (pq OrderQueue) Less(i, j int) bool {
	return pq[i].ArrivalTime < pq[j].ArrivalTime
}

func (pq OrderQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *OrderQueue) Push(x interface{}) {
	item := x.(*Order)
	*pq = append(*pq, item)
}

func (pq *OrderQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// processOrders efficiently processes orders based on their arrival time
func processOrders(orders []*Order) {
	orderQueue := make(OrderQueue, 0)
	heap.Init(&orderQueue)

	currentTime := 0

	for _, order := range orders {
		heap.Push(&orderQueue, order)
	}

	for orderQueue.Len() > 0 {
		// Pop the order with the earliest arrival time
		currentOrder := heap.Pop(&orderQueue).(*Order)

		// Wait for the order's arrival time if it's not yet arrived
		if currentOrder.ArrivalTime > currentTime {
			currentTime = currentOrder.ArrivalTime
		}

		fmt.Printf("Processing Order ID: %d, Arrival Time: %d, Processing Time: %d\n", currentOrder.ID, currentOrder.ArrivalTime, currentOrder.ProcessingTime)

		// Simulate order processing
		currentTime += currentOrder.ProcessingTime
	}
}

const (
	numOrders         = 20
	maxProcessingTime = 10
	maxArrivalTime    = 20
)

func main() {
	rand.Seed(time.Now().UnixNano())

	orders := make([]*Order, 0, numOrders)

	for i := 0; i < numOrders; i++ {
		processingTime := rand.Intn(maxProcessingTime) + 1
		arrivalTime := rand.Intn(maxArrivalTime)
		orders = append(orders, &Order{
			ID:             i,
			ProcessingTime: processingTime,
			ArrivalTime:    arrivalTime,
		})
	}

	fmt.Println("Orders before processing:")
	for _, order := range orders {
		fmt.Printf("Order ID: %d, Arrival Time: %d, Processing Time: %d\n", order.ID, order.ArrivalTime, order.ProcessingTime)
	}

	fmt.Println("\nProcessing Orders efficiently based on arrival time:")
	processOrders(orders)
}
