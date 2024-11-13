package main

import (
	"fmt"
	"sync"
	"time"
)

// Customer struct
type Customer struct {
	ID     int
	Age    int
	Gender string
}

// SegmentationResult struct
type SegmentationResult struct {
	CustomerID int
	Segment    string
}

// SegmentCustomers function segments customers based on age and gender
func SegmentCustomers(customers []Customer) []SegmentationResult {
	var results []SegmentationResult

	wg := &sync.WaitGroup{}
	c := make(chan Customer)

	// Worker goroutines that perform segmentation in parallel
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for customer := range c {
				var segment string
				if customer.Age < 18 {
					segment = "Young"
				} else if customer.Age < 30 {
					segment = "Adult"
				} else {
					segment = "Senior"
				}
				results = append(results, SegmentationResult{
					CustomerID: customer.ID,
					Segment:    segment,
				})
			}
		}()
	}

	// Send customers to worker goroutines
	for _, customer := range customers {
		c <- customer
	}

	close(c)
	wg.Wait()

	return results
}

// main function
func main() {
	// Sample customers
	customers := []Customer{
		{ID: 1, Age: 17, Gender: "M"},
		{ID: 2, Age: 25, Gender: "F"},
		{ID: 3, Age: 60, Gender: "M"},
		{ID: 4, Age: 22, Gender: "F"},
		{ID: 5, Age: 19, Gender: "M"},
		// Add more customers as needed
	}

	// Start time
	start := time.Now()

	// Segment customers with parallel processing
	results := SegmentCustomers(customers)

	// End time
	end := time.Now()

	// Display segmentation results
	fmt.Println("Segmentation Results:")
	for _, result := range results {
		fmt.Printf("Customer ID: %d, Segment: %s\n", result.CustomerID, result.Segment)
	}

	// Display execution time
	fmt.Printf("Execution Time: %s\n", end.Sub(start))
}
