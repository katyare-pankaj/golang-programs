package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add the goroutines to the WaitGroup
	wg.Add(2)

	// Start goroutines
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Goroutine 1 finished")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Goroutine 2 finished")
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines finished")
}
