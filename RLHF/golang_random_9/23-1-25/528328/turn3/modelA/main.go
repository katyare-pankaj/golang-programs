package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Product struct to represent a product in the warehouse
type Product struct {
	ID       int
	Name     string
	Quantity int
}

// Employee struct to represent each employee
type Employee struct {
	ID int
}

var products = []Product{
	{ID: 1, Name: "Laptop", Quantity: 30},
	{ID: 2, Name: "Smartphone", Quantity: 50},
	{ID: 3, Name: "Headphones", Quantity: 100},
}

// Mutex for locking product list during updates
var mutex = &sync.Mutex{}

func (e Employee) restockProduct(productID, quantity int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate restocking time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	mutex.Lock()
	defer mutex.Unlock()

	for i, product := range products {
		if product.ID == productID {
			products[i].Quantity += quantity
			fmt.Printf("EmployeeID: %d restocked ProductID: %d with Quantity: %d. New Quantity: %d\n", e.ID, productID, quantity, products[i].Quantity)
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	// Simulating different restocking tasks
	restockingTasks := []struct {
		employeeID int
		productID  int
		quantity   int
	}{
		{employeeID: 1, productID: 1, quantity: 10},
		{employeeID: 2, productID: 2, quantity: 20},
		{employeeID: 3, productID: 3, quantity: 5},
		{employeeID: 4, productID: 1, quantity: 7},
		{employeeID: 5, productID: 2, quantity: 10},
	}

	for _, task := range restockingTasks {
		employee := Employee{ID: task.employeeID}
		wg.Add(1)
		go employee.restockProduct(task.productID, task.quantity, &wg)
	}

	wg.Wait()

	fmt.Println("\nAll restocking tasks completed.")
	for _, product := range products {
		fmt.Printf("ProductID: %d, ProductName: %s, Total Quantity: %d\n", product.ID, product.Name, product.Quantity)
	}
}
