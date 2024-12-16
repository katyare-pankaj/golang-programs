package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Simulated legacy service for fetching data
func legacyDataService(wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()
	// Simulate fetching data from a legacy system
	resultChan <- "Data from Legacy System"
}

// New service to process data
func newDataService(wg *sync.WaitGroup, input string) {
	defer wg.Done()
	// Simulate processing of legacy data
	fmt.Println("Processing:", input)
	// Here, you can add further business logic
}

// Main function to transition services
func main() {
	var wg sync.WaitGroup
	resultChan := make(chan string)

	// Spin up a concurrent task for legacy data service
	wg.Add(1)
	go legacyDataService(&wg, resultChan)

	// Wait for the legacy service to send data
	go func() {
		wg.Wait()
		close(resultChan) // close the channel when done
	}()

	// Handle incoming data concurrently
	for result := range resultChan {
		wg.Add(1)
		go newDataService(&wg, result)
	}

	// Wait for all processing to complete
	wg.Wait()
	fmt.Println("All services have completed.")
}

// Dummy HTTP handler function
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the new Go application!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting...")
	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
