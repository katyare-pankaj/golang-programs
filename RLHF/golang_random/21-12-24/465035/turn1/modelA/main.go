package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated IoT data collection and processing
func collectData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Data collection started.")
	time.Sleep(time.Second * 2) // Simulate data collection delay
	fmt.Println("Data collection completed.")
}

func processData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Data processing started.")
	time.Sleep(time.Second * 1) // Simulate data processing delay
	fmt.Println("Data processing completed.")
}

func sendData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Data sending started.")
	time.Sleep(time.Second * 1) // Simulate data sending delay
	fmt.Println("Data sending completed.")
}

func main() {
	var wg sync.WaitGroup

	// Start data collection, processing, and sending in parallel
	wg.Add(1) // Increment the waitgroup for data collection
	go collectData(&wg)

	wg.Add(1) // Increment the waitgroup for data processing
	go processData(&wg)

	wg.Add(1) // Increment the waitgroup for data sending
	go sendData(&wg)

	// Main thread waits for all worker threads to finish
	fmt.Println("Main thread waiting for all tasks to complete...")
	wg.Wait()

	fmt.Println("All tasks completed. Main thread exiting.")
}
