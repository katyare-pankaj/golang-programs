package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Product represents an item in the inventory
type Product struct {
	ID       int
	Name     string
	Quantity int
}

// Order represents a user's order
type Order struct {
	UserID    int
	ProductID int
	Quantity  int
}

// Inventory holds a list of products and manages stock concurrently
type Inventory struct {
	products map[int]*Product
	mutex    sync.Mutex
}

// NewInventory initializes the inventory with some products
func NewInventory(products []Product) *Inventory {
	productMap := make(map[int]*Product, len(products))
	for i := range products {
		productMap[products[i].ID] = &products[i]
	}
	return &Inventory{products: productMap}
}

// ProcessOrder processes the order for a product, reducing the quantity in stock
func (inv *Inventory) ProcessOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	inv.mutex.Lock()
	defer inv.mutex.Unlock()

	product, ok := inv.products[order.ProductID]
	if ok {
		if product.Quantity >= order.Quantity {
			product.Quantity -= order.Quantity
			fmt.Printf("Order processed for UserID %d: %d of %s\n", order.UserID, order.Quantity, product.Name)
		} else {
			fmt.Printf("Insufficient stock for UserID %d's order: %d of %s\n", order.UserID, order.Quantity, product.Name)
		}
	} else {
		fmt.Printf("Product not found for UserID %d's order\n", order.UserID)
	}
}

func main() {
	products := []Product{
		{ID: 1, Name: "Laptop", Quantity: 10},
		{ID: 2, Name: "Smartphone", Quantity: 20},
		{ID: 3, Name: "Tablet", Quantity: 15},
	}

	orders := []Order{
		{UserID: 1, ProductID: 1, Quantity: 2},
		{UserID: 2, ProductID: 1, Quantity: 5},
		{UserID: 3, ProductID: 2, Quantity: 1},
		{UserID: 4, ProductID: 3, Quantity: 5},
		{UserID: 5, ProductID: 3, Quantity: 10},
		{UserID: 6, ProductID: 2, Quantity: 22},
	}

	inventory := NewInventory(products)

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	for _, order := range orders {
		wg.Add(1)
		go func(order Order) {
			// Simulate varied processing time
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			inventory.ProcessOrder(order, &wg)
		}(order)
	}

	wg.Wait()
	fmt.Println("All orders have been processed. Final inventory status:")

	for _, product := range inventory.products {
		fmt.Printf("Product: %s, Remaining Quantity: %d\n", product.Name, product.Quantity)
	}
}
