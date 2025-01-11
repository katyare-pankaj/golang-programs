package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"Apple", "Banana", "Cherry", "Orange", "Pineapple"}
	processFruits(fruits)
}

func processFruits(fruits []string) {
	var wg sync.WaitGroup
	wg.Add(len(fruits))

	// Safely use range in a goroutine by passing a copy of the fruit variable
	for _, fruit := range fruits {
		go func(fruit string) {
			defer wg.Done()
			fmt.Println("Processing fruit:", fruit)
			// Simulate some work
			work()
		}(fruit)
	}

	wg.Wait()
	fmt.Println("All fruits processed.")
}

func work() {
	// Simulate work using a delay
}
