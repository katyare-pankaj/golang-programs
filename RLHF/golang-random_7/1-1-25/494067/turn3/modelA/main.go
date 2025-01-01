package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Item struct {
	Name     string
	Quantity int
}

func main() {
	// Initialize a sync.Map to store item inventories
	inventory := sync.Map{}

	// Random items for our inventory
	items := []string{"apple", "banana", "orange", "bread", "milk", "cheese", "chocolate"}

	// Number of items to track
	numItems := 20

	// Number of goroutines for concurrent operations
	numGoroutines := 10

	// Initialize inventory with random items and quantities
	for i := 0; i < numItems; i++ {
		item := items[i%len(items)]
		quantity := rand.Intn(100)
		inventory.Store(item, &Item{Name: item, Quantity: quantity})
	}

	// Create a WaitGroup to coordinate goroutines
	var wg sync.WaitGroup

	// Goroutine to perform operations (add, remove, query) on the inventory
	performOperation := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			operation := rand.Intn(3)
			switch operation {
			case 0:
				// Add a new item or increase the quantity of an existing one
				item := items[rand.Intn(len(items))]
				current, _ := inventory.LoadOrStore(item, &Item{Name: item, Quantity: 0})
				current.(*Item).Quantity += rand.Intn(50)
			case 1:
				// Remove an item or decrease the quantity of an existing one
				item := items[rand.Intn(len(items))]
				current, ok := inventory.Load(item)
				if ok {
					current.(*Item).Quantity -= rand.Intn(current.(*Item).Quantity + 1)
					if current.(*Item).Quantity <= 0 {
						inventory.Delete(item)
					} else {
						inventory.Store(item, current)
					}
				}
			default:
				// Query an item's quantity
				item := items[rand.Intn(len(items))]
				current, ok := inventory.Load(item)
				if ok {
					fmt.Printf("Current quantity of %s: %d\n", current.(*Item).Name, current.(*Item).Quantity)
				}
			}
		}
	}

	// Start the specified number of goroutines to perform operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go performOperation()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Display the final state of the inventory
	fmt.Println("Final Inventory:")
	inventory.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", value.(*Item).Name, value.(*Item).Quantity)
		return true
	})
}
