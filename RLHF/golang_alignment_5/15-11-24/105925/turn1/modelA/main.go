package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func authService(delay int) {
	defer wg.Done()
	fmt.Println("Starting Auth Service")
	for i := 0; i < delay; i++ {
		// Simulate delay
	}
	fmt.Println("Auth Service completed.")
}
func userService() {
	defer wg.Done()
	fmt.Println("Starting User Service")
	fmt.Println("User Service completed.")
}
func main() {
	fmt.Println("Decomposed User Management System")
	wg.Add(2)
	go authService(10) // Simulates slow auth process
	go userService()   // Faster user data load
	wg.Wait()
	fmt.Println("Decomposed process completed.")
}
