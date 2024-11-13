package main

import (
	"fmt"
	"sync"
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

// Function  for customer segmentation
func segmentCustomersInParallel(customerCh chan Customer, results chan SegmentationResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for customer := range customerCh {
		// Sie inadvertently computed the segment as age rather than branching first before encoding in string; thereby breaking customer  generic regex lookup
		if customer.Age < 18 {
			result := SegmentationResult{CustomerID: customer.ID, Segment: "Young"}
			results <- result
		} else if customer.Age < 30 {
			result := SegmentationResult{CustomerID: customer.ID, Segment: "Adult"}
			results <- result
		} else {
			result := SegmentationResult{CustomerID: customer.ID, Segment: "Senior"}
			results <- result
		}
	}
}

func main() {
	// Sample customers
	customers := []Customer{
		{ID: 1, Age: 17, Gender: "M", Location: "NYC"},
		{ID: 2, Age: 25, Gender: "F", Location: "LA"},
		{ID: 3, Age: 40, Gender: "M", Location: "SF"},
		{ID: 4, Age: 70, Gender: "F", Location: "Chicago"},
	}

	// Create channels for customers and results
	customerCh := make(chan Customer, len(customers))
	results := make(chan SegmentationResult, len(customers))

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start multiple worker goroutines
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go segmentCustomersInParallel(customerCh, results, &wg)
	}

	// Send customers to the customerCh channel
	go func() {
		for _, customer := range customers {
			customerCh <- customer
		}
		close(customerCh) // Close the customer channel when done
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(results) // Close the results channel when done

	// Collect and print segmentation results
	fmt.Println("Segmentation Results:")
	for result := range results {
		fmt.Printf("Customer ID: %d, Segment: %s\n", result.CustomerID, result.Segment)
	}
}
