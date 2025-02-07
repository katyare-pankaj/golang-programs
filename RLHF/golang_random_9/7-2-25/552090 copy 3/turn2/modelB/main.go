package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Distance int
}

func processOrder(order Order, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Processing Order ID: %d, Distance: %d km\n", order.ID, order.Distance)
	time.Sleep(time.Duration(order.Distance) * time.Millisecond)
	fmt.Printf("Order ID: %d completed.\n", order.ID)
}

func main() {
	orders := []Order{
		{ID: 1, Distance: 50},
		{ID: 2, Distance: 75},
		{ID: 3, Distance: 30},
		{ID: 4, Distance: 40},
		{ID: 5, Distance: 60},
	}

	var wg sync.WaitGroup

	wg.Add(len(orders))

	for _, order := range orders {
		go processOrder(order, &wg)
	}

	wg.Wait()
	fmt.Println("All orders processed.")
}
