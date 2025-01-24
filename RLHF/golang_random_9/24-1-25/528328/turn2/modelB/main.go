package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate a product in the e-commerce platform
type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Simulate a User order
type Order struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// User struct to represent each user
type User struct {
	UserID int
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
	{UserID: 6, ProductID: 3, Quantity: 2},
	{UserID: 7, ProductID: 2, Quantity: 4},
	{UserID: 8, ProductID: 1, Quantity: 3},
	{UserID: 9, ProductID: 3, Quantity: 1},
	{UserID: 10, ProductID: 1, Quantity: 2},
}

// Total number of orders processed successfully
var totalOrdersProcessed = 0

// Mutex for locking product list during updates
var mutex = &sync.Mutex{}

func (u User) placeOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate processing time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	mutex.Lock()
	defer mutex.Unlock()

	for i, product := range products {
		if product.ID == order.ProductID {
			if product.Quantity >= order.Quantity {
				products[i].Quantity -= order.Quantity
				fmt.Printf("UserID: %d placed an order for ProductID: %d, Quantity: %d. Updated Quantity: %d\n", u.UserID, order.ProductID, order.Quantity, products[i].Quantity)
				totalOrdersProcessed++
			} else {
				fmt.Printf("UserID: %d couldn't place an order for ProductID: %d due to insufficient stock. Remaining Quantity: %d\n", u.UserID, order.ProductID, products[i].Quantity)
			}
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for _, order := range orders {
		user := User{UserID: order.UserID}
		wg.Add(1)
		go user.placeOrder(order, &wg)
	}

	wg.Wait()

	fmt.Println("\nAll orders processed.")
	fmt.Println("Total orders processed successfully:", totalOrdersProcessed)
	for _, product := range products {
		fmt.Printf("ProductID: %d, ProductName: %s, Remaining Quantity: %d\n", product.ID, product.Name, product.Quantity)
	}
}
