package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Immutable data structure to represent a product's price and quantity
type Product struct {
	Price    float64
	Quantity int
}

// Lazy-evaluated data structure to hold a inventory of products
type Inventory struct {
	mu           sync.Mutex
	products     map[string]Product
	lazyData     *lazyData
	lazyDataCalc sync.Once
}

type lazyData struct {
	totalValue float64
}

// Create a new Inventory instance
func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]Product),
	}
}

// Add a new product or update its quantity to the Inventory
func (i *Inventory) UpdateProduct(product string, price float64, quantity int) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.products[product] = Product{Price: price, Quantity: quantity}
	i.lazyDataCalc.Do(func() {
		i.lazyData = &lazyData{}
		for _, p := range i.products {
			i.lazyData.totalValue += p.Price * float64(p.Quantity)
		}
	})
	i.lazyData = nil // Invalidate the lazy data when updating a product
}

// Get the total value of the inventory (lazy evaluation)
func (i *Inventory) TotalValue() float64 {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.lazyDataCalc.Do(func() {
		i.lazyData = &lazyData{}
		for _, p := range i.products {
			i.lazyData.totalValue += p.Price * float64(p.Quantity)
		}
	})
	return i.lazyData.totalValue
}

// Simulate real-time data processing by updating inventory products
func simulateRealTimeDataProcessing(i *Inventory) {
	rand.Seed(time.Now().UnixNano())
	products := []string{"ProductA", "ProductB", "ProductC"}

	for {
		product := products[rand.Intn(len(products))]
		price := rand.Float64() * 100
		quantity := rand.Intn(100)

		i.UpdateProduct(product, price, quantity)

		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	inv := NewInventory()

	// Start the simulation of real-time data processing
	go simulateRealTimeDataProcessing(inv)

	// Example usage:
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("Total Inventory Value: $%.2f\n", inv.TotalValue())
	}
}
