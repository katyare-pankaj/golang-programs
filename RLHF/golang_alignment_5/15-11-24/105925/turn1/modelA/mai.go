package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func simulateWork(operation string, delay int) {
	defer wg.Done()
	fmt.Println("Starting", operation)
	for i := 0; i < delay; i++ {
		// Delay for demonstration purposes
	}
	fmt.Println("Completed", operation)
}
func main() {
	fmt.Println("Monolithic User Management System")
	wg.Add(2)
	go simulateWork("User Login", 10)       // Simulates slow auth process
	go simulateWork("User Profile Load", 2) // Simulates faster user data load
	wg.Wait()
	fmt.Println("Monolithic process completed.")
}
