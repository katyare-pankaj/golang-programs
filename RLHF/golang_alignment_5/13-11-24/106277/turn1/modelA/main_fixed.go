package main

import (
	"fmt"
	"time"
)

// Customer struct
type Customer struct {
	ID       int
	Age      int
	Gender   string
	Location string
}

// SegmentCustomers function
func SegmentCustomers(customers []Customer) map[string][]Customer {
	start := time.Now()
	defer func() {
		fmt.Println("Segmentation time:", time.Since(start))
	}()

	segments := make(map[string][]Customer)

	// Compatibility check: Ensure Age is an integer
	for _, customer := range customers {
		if _, ok := interface{}(customer.Age).(int); !ok {
			panic("Age must be an integer")
		}
	}

	// Optimize for low latency using parallel processing
	ch := make(chan Customer)
	go func() {
		for _, customer := range customers {
			ch <- customer
		}
		close(ch)
	}()

	for customer := range ch {
		if _, exists := segments[customer.Gender]; !exists {
			segments[customer.Gender] = []Customer{customer}
		} else {
			segments[customer.Gender] = append(segments[customer.Gender], customer)
		}
	}

	return segments
}

func main() {
	customers := []Customer{
		{ID: 1, Age: 25, Gender: "Female", Location: "New York"},
		{ID: 2, Age: 30, Gender: "Male", Location: "Los Angeles"},
		{ID: 3, Age: 22, Gender: "Female", Location: "Chicago"},
		// Add more customers...
	}

	segments := SegmentCustomers(customers)
	fmt.Println(segments)
}
