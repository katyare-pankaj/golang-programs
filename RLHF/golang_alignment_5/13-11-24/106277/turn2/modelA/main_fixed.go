package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Customer struct
type Customer struct {
	ID       int
	Age      int
	Gender   string
	Location string
}

// SegmentationResult struct
type SegmentationResult struct {
	CustomerID int
	Segment    string
}

// SegmentCustomer function with error handling
func SegmentCustomer(ctx context.Context, customer Customer, results chan<- SegmentationResult) error {
	var segment string
	if customer.Age < 18 {
		segment = "Young"
	} else if customer.Age < 30 {
		segment = "Adult"
	} else {
		segment = "Senior"
	}

	select {
	case results <- SegmentationResult{CustomerID: customer.ID, Segment: segment}:
		fmt.Printf("Segmented Customer ID: %d, Segment: %s\n", customer.ID, segment) // Debug print
		return nil
	case <-ctx.Done():
		return fmt.Errorf("context cancelled: %w", ctx.Err())
	}
}

// SegmentCustomersConcurrently function to process customers in parallel
func SegmentCustomersConcurrently(customers []Customer) []SegmentationResult {
	const numWorkers = 4

	customersChan := make(chan Customer, len(customers))
	var results []SegmentationResult
	wg := &sync.WaitGroup{}

	// Use a channel to send the segmentation results
	resultsChan := make(chan SegmentationResult, len(customers))

	// Create a context with a timeout to handle slow customers
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for customer := range customersChan {
				// Call SegmentCustomer for each customer
				if err := SegmentCustomer(ctx, customer, resultsChan); err != nil {
					fmt.Println("error segmenting customer:", err)
				}
			}
		}()
	}

	// Create a channel for customers to send to workers

	// Send customers to the channel
	go func() {
		for _, customer := range customers {
			customersChan <- customer
		}
		close(customersChan)
	}()

	// Wait for all worker goroutines to complete
	wg.Wait()

	// Close the results channel once all workers are done
	close(resultsChan)

	// Collect results from the results channel
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func main() {
	// Sample customers
	customers := []Customer{
		{ID: 1, Age: 17, Gender: "M", Location: "NYC"},
		{ID: 2, Age: 25, Gender: "F", Location: "LA"},
		{ID: 3, Age: 40, Gender: "M", Location: "SF"},
		{ID: 4, Age: 70, Gender: "F", Location: "Chicago"},
		// Add more customers as needed
	}

	start := time.Now()
	results := SegmentCustomersConcurrently(customers)
	end := time.Now()

	// Print segmentation results
	fmt.Println("Segmentation Results:")
	for _, result := range results {
		fmt.Printf("Customer ID: %d, Segment: %s\n", result.CustomerID, result.Segment)
	}
	fmt.Printf("Execution Time: %s\n", end.Sub(start))
}
