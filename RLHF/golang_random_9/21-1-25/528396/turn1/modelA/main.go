package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DataItem represents a single unit of data to be migrated.
type DataItem struct {
	ID    int
	Value string
}

// migrateData simulates the migration of a single DataItem.
// It returns an error if the migration fails.
func migrateData(item DataItem) error {
	// Simulate a random chance of failure
	if rand.Intn(100) < 10 { // 10% chance of failure
		return fmt.Errorf("failed to migrate item ID %d", item.ID)
	}

	// Simulate some processing time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

	fmt.Printf("Successfully migrated item ID %d\n", item.ID)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Define some data items to migrate
	dataItems := []DataItem{
		{ID: 1, Value: "Data1"},
		{ID: 2, Value: "Data2"},
		{ID: 3, Value: "Data3"},
		{ID: 4, Value: "Data4"},
		{ID: 5, Value: "Data5"},
		{ID: 6, Value: "Data6"},
		{ID: 7, Value: "Data7"},
		{ID: 8, Value: "Data8"},
		{ID: 9, Value: "Data9"},
		{ID: 10, Value: "Data10"},
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make([]error, 0)

	for _, item := range dataItems {
		wg.Add(1)

		// Launch a goroutine for each data item
		go func(item DataItem) {
			defer wg.Done()

			if err := migrateData(item); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(item)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Handle any errors that occurred during migration
	if len(errors) > 0 {
		fmt.Println("The following errors occurred during migration:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("All data items migrated successfully.")
	}
}
