package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Count the number of goroutines to wait for
	numGoroutines := 3
	wg.Add(numGoroutines)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 started...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 finished.")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 started...")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 finished.")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3 started...")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine 3 finished.")
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines completed. Main program continuing...")
	// Further execution of the main program can go here.
	fmt.Println("Main program finished.")
}
