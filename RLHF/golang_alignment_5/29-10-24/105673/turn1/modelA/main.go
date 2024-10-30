package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// Order represents a supply chain order
type Order struct {
	ID       int
	Priority int // Higher value means higher priority
	Size     int
}

// Queue represents a queue of orders
type Queue struct {
	Orders    []*Order
	totalSize int
}

// PriorityQueue implements a priority queue for Orders
type PriorityQueue []*Order

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority means smaller value in the priority queue
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Order)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// AddOrder adds an order to the appropriate queue
func AddOrder(orders map[int]*Queue, pq *PriorityQueue, newOrder *Order) {
	// Find the queue with the least total size or create a new one if all queues are full
	minQueue := findMinQueue(orders)
	if minQueue == nil {
		minQueue = &Queue{}
		orders[len(orders)] = minQueue
	}

	minQueue.Orders = append(minQueue.Orders, newOrder)
	minQueue.totalSize += newOrder.Size

	// Add the order to the priority queue for dynamic reallocation
	heap.Push(pq, newOrder)
}

func findMinQueue(queues map[int]*Queue) *Queue {
	var minQueue *Queue
	minSize := intMax

	for _, queue := range queues {
		if queue.totalSize < minSize {
			minQueue = queue
			minSize = queue.totalSize
		}
	}
	return minQueue
}

const (
	intMax       = int(^uint(0) >> 1)
	numQueues    = 3
	maxOrderSize = 10
	numOrders    = 20
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ordersMap := make(map[int]*Queue)
	var priorityQueue PriorityQueue
	heap.Init(&priorityQueue)

	// Generate random orders
	for i := 0; i < numOrders; i++ {
		orderSize := rand.Intn(maxOrderSize) + 1
		newOrder := &Order{
			ID:       i,
			Priority: rand.Intn(10), // Random priority between 0 and 9
			Size:     orderSize,
		}
		AddOrder(ordersMap, &priorityQueue, newOrder)
	}

	// Process orders from priority queue
	for priorityQueue.Len() > 0 {
		currentOrder := heap.Pop(&priorityQueue).(*Order)
		fmt.Printf("Processing Order ID: %d, Priority: %d, Size: %d\n", currentOrder.ID, currentOrder.Priority, currentOrder.Size)

		// Update queue sizes after processing
		for _, queue := range ordersMap {
			for i, order := range queue.Orders {
				if order.ID == currentOrder.ID {
					queue.Orders = append(queue.Orders[:i], queue.Orders[i+1:]...)
					queue.totalSize -= currentOrder.Size
					break
				}
			}
		}
	}

	fmt.Println("Order processing completed.")
}
