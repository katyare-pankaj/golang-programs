package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Immutable data structure to represent a customer's order
type Order struct {
	ID       int
	Customer string
	Items    []string
}

// Lazy-evaluated data structure to hold a collection of orders
type OrderBook struct {
	mu       sync.Mutex
	orders   map[int]Order
	lazyData *lazyData
}

type lazyData struct {
	customers map[string]int
}

// Create a new OrderBook instance
func NewOrderBook() *OrderBook {
	return &OrderBook{
		orders:   make(map[int]Order),
		lazyData: &lazyData{},
	}
}

// Add a new order to the OrderBook
func (ob *OrderBook) AddOrder(order Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	ob.orders[order.ID] = order
	ob.lazyData = nil // Invalidate the lazy data when adding a new order
}

// Get the order by its ID
func (ob *OrderBook) GetOrder(id int) (Order, bool) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	order, ok := ob.orders[id]
	return order, ok
}

// Get a count of orders placed by a specific customer (lazy evaluation)
func (ob *OrderBook) GetCustomerOrderCount(customer string) int {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	// Lazy initialization of lazyData
	if ob.lazyData == nil {
		ob.lazyData = &lazyData{
			customers: make(map[string]int),
		}
		for _, order := range ob.orders {
			ob.lazyData.customers[order.Customer]++
		}
	}

	return ob.lazyData.customers[customer]
}

// Simulate real-time data processing by adding random orders to the OrderBook
func simulateRealTimeDataProcessing(ob *OrderBook) {
	rand.Seed(time.Now().UnixNano())

	for {
		orderID := rand.Intn(1000)
		customer := fmt.Sprintf("Customer%d", rand.Intn(100))
		items := []string{"Item1", "Item2", "Item3"}
		rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

		order := Order{ID: orderID, Customer: customer, Items: items}
		ob.AddOrder(order)

		time.Sleep(time.Millisecond * 50) // Simulate real-time processing delay
	}
}

func main() {
	ob := NewOrderBook()

	// Start the simulation of real-time data processing
	go simulateRealTimeDataProcessing(ob)

	// Example usage:
	for {
		time.Sleep(time.Second * 1)

		customer := "Customer42"
		orderCount := ob.GetCustomerOrderCount(customer)
		fmt.Printf("Customer %s has placed %d orders.\n", customer, orderCount)

		// Example of deep equality comparison for immutable data structures
		order1 := Order{ID: 1, Customer: "CustomerX", Items: []string{"ItemA", "ItemB"}}
		order2 := Order{ID: 1, Customer: "CustomerX", Items: []string{"ItemA", "ItemB"}}
		if cmp.Equal(order1, order2) {
			fmt.Println("Orders are equal.")
		} else {
			fmt.Println("Orders are different.")
		}
	}
}
