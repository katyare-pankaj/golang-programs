package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(wg *sync.WaitGroup, message string, delay time.Duration) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup

	// Number of goroutines to synchronize
	numGoroutines := 5

	// Start multiple goroutines
	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		message := fmt.Sprintf("Message %d", i)
		delay := time.Duration(i) * time.Second
		go printMessage(&wg, message, delay)
	}

	// Wait for all goroutines to complete their tasks
	wg.Wait()

	fmt.Println("All goroutines have finished printing messages.")
}
