package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Product struct
type Product struct {
	ID   int
	Name string
	Qty  int
}

// Employee struct
type Employee struct {
	ID int
}

var products = []Product{
	{ID: 1, Name: "Laptop", Qty: 20},
	{ID: 2, Name: "Mouse", Qty: 50},
	{ID: 3, Name: "Keyboard", Qty: 30},
	{ID: 4, Name: "Monitor", Qty: 40},
	{ID: 5, Name: "CPU", Qty: 10},
}

var employees = []Employee{
	{ID: 1},
	{ID: 2},
	{ID: 3},
	{ID: 4},
}

// Restock product function
func (e Employee) restock(product *Product, restockQty int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate restocking time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Printf("Employee %d is restocking Product %d: %s, Quantity: %d\n", e.ID, product.ID, product.Name, restockQty)
	product.Qty += restockQty
}

var mutex sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// Randomly assign restock tasks to employees
	for i := 0; i < 20; i++ {
		product := products[rand.Intn(len(products))]
		employee := employees[rand.Intn(len(employees))]
		restockQty := rand.Intn(20) + 1 // Restock quantity between 1 and 20

		wg.Add(1)
		go employee.restock(&product, restockQty, &wg)
	}

	wg.Wait()

	fmt.Println("\nRestocking complete.")
	fmt.Println("Updated Product Inventory:")
	for _, product := range products {
		fmt.Printf("ProductID: %d, ProductName: %s, Remaining Quantity: %d\n", product.ID, product.Name, product.Qty)
	}
}
