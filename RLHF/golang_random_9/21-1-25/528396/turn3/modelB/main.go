package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Transaction represents a single financial transaction.
type Transaction struct {
	ID     int
	Amount float64
}

// simulateTransactionProcessing simulates the processing of a single Transaction.
// It returns an error if the processing fails.
func simulateTransactionProcessing(transaction Transaction) error {
	// Simulate a random processing time (0 to 500 milliseconds)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

	// Simulate a 10% chance of failure
	if rand.Intn(10) == 0 {
		return fmt.Errorf("transaction ID %d processing failed", transaction.ID)
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Define a slice of transactions to process
	transactions := []Transaction{
		{ID: 1, Amount: 100.00},
		{ID: 2, Amount: 50.00},
		{ID: 3, Amount: 75.00},
		{ID: 4, Amount: 20.00},
		{ID: 5, Amount: 90.00},
		// Add more transactions as needed
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalProcessed := 0.0
	errors := make([]error, 0)

	// Launch goroutines to process each transaction concurrently
	for _, transaction := range transactions {
		wg.Add(1)

		go func(transaction Transaction) {
			defer wg.Done()

			if err := simulateTransactionProcessing(transaction); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			} else {
				mu.Lock()
				totalProcessed += transaction.Amount
				mu.Unlock()
			}
		}(transaction)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Report any errors encountered during processing
	if len(errors) > 0 {
		fmt.Println("The following errors occurred during transaction processing:")
		for _, err := range errors {
			fmt.Println(err)
		}
	}

	// Display the total amount processed
	fmt.Printf("Total amount processed: %.2f\n", totalProcessed)
}
