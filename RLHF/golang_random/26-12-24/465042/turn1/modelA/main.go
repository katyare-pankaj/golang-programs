package main

import (
	"fmt"
	"sync"
	"time"
)

func processData(id int, data []int, wg *sync.WaitGroup) {
	defer wg.Done()                                   // Mark the routine as complete
	time.Sleep(time.Duration(100) * time.Millisecond) // Simulate processing time
	fmt.Printf("Processed data for id: %d\n", id)
}

func main() {
	var wg sync.WaitGroup // Initialize WaitGroup
	var data [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Add the main routine to the WaitGroup
	wg.Add(1)

	// Start processing in concurrent goroutines
	for i, d := range data {
		wg.Add(1) // Add each goroutine to the WaitGroup
		go processData(i, d, &wg)
	}

	// Wait for all goroutines to complete before printing a final message
	wg.Wait()
	fmt.Println("All data processing complete.")
}
