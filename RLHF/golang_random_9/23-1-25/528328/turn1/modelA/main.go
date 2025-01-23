package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate a product in the e-commerce platform
type Product struct {
	ID       int
	Name     string
	Quantity int
}

// Simulate a User order
type Order struct {
	UserID    int
	ProductID int
	Quantity  int
}

var products = []Product{
	{ID: 1, Name: "Laptop", Quantity: 30},
	{ID: 2, Name: "Smartphone", Quantity: 50},
	{ID: 3, Name: "Headphones", Quantity: 100},
}

var orders = []Order{
	{UserID: 1, ProductID: 1, Quantity: 2},
	{UserID: 2, ProductID: 2, Quantity: 1},
	{UserID: 3, ProductID: 3, Quantity: 5},
	{UserID: 4, ProductID: 2, Quantity: 3},
	{UserID: 5, ProductID: 1, Quantity: 1},
}

// Mutex for locking product list during updates
var mutex = &sync.Mutex{}

func processOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	mutex.Lock()
	defer mutex.Unlock()
	for i, product := range products {
		if product.ID == order.ProductID {
			if product.Quantity >= order.Quantity {
				products[i].Quantity -= order.Quantity
				fmt.Printf("Order processed for UserID: %d, ProductID: %d, Quantity: %d\n", order.UserID, order.ProductID, order.Quantity)
			} else {
				fmt.Printf("Couldn't process order for UserID: %d due to insufficient stock for ProductID: %d\n", order.UserID, order.ProductID)
			}
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(1)
		go processOrder(order, &wg)
	}

	wg.Wait()

	fmt.Println("All orders processed")
	for _, product := range products {
		fmt.Printf("ProductID: %d, ProductName: %s, Remaining Quantity: %d\n", product.ID, product.Name, product.Quantity)
	}
}
