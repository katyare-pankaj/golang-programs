package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataItem struct {
	ID    int
	Value string
}

type MigrationResult struct {
	Item DataItem
	Err  error
}

func migrateData(item DataItem, resultChan chan<- MigrationResult) {
	defer close(resultChan)
	err := SimulateMigration(item)
	resultChan <- MigrationResult{Item: item, Err: err}
}

func SimulateMigration(item DataItem) error {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	if rand.Intn(100) < 10 {
		return errors.New("migration failed")
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dataItems := []DataItem{
		{ID: 1, Value: "Data1"},
		{ID: 2, Value: "Data2"},
		// Add more data items as needed
	}

	var wg sync.WaitGroup
	numGoroutines := 10 // Adjust the number of goroutines as needed for concurrency
	resultChan := make(chan MigrationResult)

	for _, item := range dataItems {
		wg.Add(1)
		go func(item DataItem) {
			defer wg.Done()
			migrateData(item, resultChan)
		}(item)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	migrationSuccesses := 0
	var migrationErrors []error

	for result := range resultChan {
		if result.Err != nil {
			migrationErrors = append(migrationErrors, fmt.Errorf("failed to migrate item ID %d: %w", result.Item.ID, result.Err))
		} else {
			migrationSuccesses++
		}
	}

	if len(migrationErrors) == 0 {
		fmt.Printf("All %d data items migrated successfully.\n", migrationSuccesses)
	} else {
		fmt.Printf("%d data items migrated successfully, but the following %d errors occurred:\n", migrationSuccesses, len(migrationErrors))
		for _, err := range migrationErrors {
			fmt.Println(err)
		}
	}
}
