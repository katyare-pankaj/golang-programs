package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Record represents a single unit of data to be migrated.
type Record struct {
	ID    int
	Value string
}

// simulateMigration simulates the processing of migrating a single Record.
// It returns an error if the migration fails.
func simulateMigration(record Record) error {
	// Simulate a random migration failure
	if rand.Intn(100) < 10 { // 10% chance of failure
		return fmt.Errorf("migration failed for record ID %d", record.ID)
	}
	// Simulate a random processing time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Printf("Migrated record ID %d successfully.\n", record.ID)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator for simulated behavior

	// Sample data records to migrate
	records := []Record{
		{ID: 1, Value: "Data1"},
		{ID: 2, Value: "Data2"},
		{ID: 3, Value: "Data3"},
		{ID: 4, Value: "Data4"},
		{ID: 5, Value: "Data5"},
		// Add more records as needed
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make([]error, 0)

	// Iterate over each record and process them concurrently
	for _, record := range records {
		wg.Add(1) // Increment the WaitGroup counter

		// Launch a goroutine to handle each record
		go func(record Record) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			if err := simulateMigration(record); err != nil {
				mu.Lock()
				errors = append(errors, err) // Safely handle errors using a mutex
				mu.Unlock()
			}
		}(record)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check and report any errors encountered during the migration
	if len(errors) > 0 {
		log.Println("Some migrations failed:")
		for _, err := range errors {
			log.Println(err)
		}
	} else {
		fmt.Println("All records migrated successfully.")
	}
}
