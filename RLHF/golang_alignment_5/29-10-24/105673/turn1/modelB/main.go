package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Order struct {
	ID       int
	Priority int
}

type Queue struct {
	Orders []Order
}

type SupplyChainManager struct {
	Queues       []Queue
	MaxQueueSize int
}

func (scm *SupplyChainManager) AssignOrder(order Order) {
	// Find the queue with the lowest priority
	queueIndex := 0
	minPriority := scm.Queues[0].Orders[0].Priority
	for i, queue := range scm.Queues {
		if len(queue.Orders) > 0 {
			if queue.Orders[0].Priority < minPriority {
				queueIndex = i
				minPriority = queue.Orders[0].Priority
			}
		}
	}

	// If the queue is full, create a new one
	if len(scm.Queues[queueIndex].Orders) == scm.MaxQueueSize {
		scm.Queues = append(scm.Queues, Queue{})
		queueIndex = len(scm.Queues) - 1
	}

	// Assign the order to the queue
	scm.Queues[queueIndex].Orders = append(scm.Queues[queueIndex].Orders, order)
	sort.Slice(scm.Queues[queueIndex].Orders, func(i, j int) bool {
		return scm.Queues[queueIndex].Orders[i].Priority < scm.Queues[queueIndex].Orders[j].Priority
	})
}

func (scm *SupplyChainManager) ProcessOrders() {
	for _, queue := range scm.Queues {
		for _, order := range queue.Orders {
			fmt.Printf("Processing order %d with priority %d\n", order.ID, order.Priority)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		queue.Orders = []Order{}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	scm := &SupplyChainManager{
		Queues:       []Queue{Queue{}},
		MaxQueueSize: 3,
	}

	// Generate random orders with priorities
	for i := 0; i < 15; i++ {
		order := Order{
			ID:       i,
			Priority: rand.Intn(10) + 1,
		}
		scm.AssignOrder(order)
	}

	fmt.Println("Initial Queue Assignment:")
	for i, queue := range scm.Queues {
		fmt.Printf("Queue %d: %v\n", i, queue.Orders)
	}

	fmt.Println("\nProcessing Orders:")
	scm.ProcessOrders()
}
