package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a database connection using a mutex
var (
	dbMutex   sync.Mutex
	dbCounter int
)

func dbOperation() {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	dbCounter++
	time.Sleep(10 * time.Millisecond) // Simulate database operation time
}

// Simulate concurrent database access using a waitgroup
func dbOperations(wg *sync.WaitGroup, numOperations int) {
	defer wg.Done()
	for i := 0; i < numOperations; i++ {
		dbOperation()
	}
}
func main() {
	startTime := time.Now()
	// Simulate 100 concurrent database operations using a mutex
	const numOperations = 100
	for i := 0; i < numOperations; i++ {
		go dbOperation()
	}
	// Wait for all operations to complete using a mutex
	time.Sleep(time.Second)
	fmt.Println("Mutex:")
	fmt.Println("Database operations:", dbCounter)
	fmt.Println("Execution time:", time.Since(startTime))
	startTime = time.Now()
	// Simulate 100 concurrent database operations using a waitgroup
	var wg sync.WaitGroup
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go dbOperations(&wg, 1)
	}
	wg.Wait()
	fmt.Println("WaitGroup:")
	fmt.Println("Database operations:", dbCounter)
	fmt.Println("Execution time:", time.Since(startTime))
}
