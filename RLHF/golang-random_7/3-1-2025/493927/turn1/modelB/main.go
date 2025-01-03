package main

import (
	"fmt"
	"sync"
)

// Future represents the result of an asynchronous operation.
type Future interface {
	Get() (interface{}, error)
}

// MigrateFunc is a function that performs data migration.
type MigrateFunc func() (interface{}, error)

// NewFuture creates a new Future from a MigrateFunc.
func NewFuture(migrate MigrateFunc) Future {
	future := &futureImpl{
		migrate: migrate,
		result:  make(chan result, 1),
	}
	go future.start()
	return future
}

type futureImpl struct {
	migrate MigrateFunc
	result  chan result
}

type result struct {
	value interface{}
	err   error
}

func (f *futureImpl) start() {
	defer close(f.result)
	value, err := f.migrate()
	f.result <- result{value, err}
}

// Get retrieves the result of the migration operation.
func (f *futureImpl) Get() (interface{}, error) {
	r := <-f.result
	return r.value, r.err
}

// Migrate performs data migration using a goroutine pool.
func Migrate(migrateFuncs []MigrateFunc, concurrency int) []Future {
	futures := make([]Future, len(migrateFuncs))
	var wg sync.WaitGroup
	wg.Add(len(migrateFuncs))

	for i, migrate := range migrateFuncs {
		go func(i int, migrate MigrateFunc) {
			defer wg.Done()
			futures[i] = NewFuture(migrate)
		}(i, migrate)
	}

	wg.Wait()
	return futures
}

func main() {
	// Example data migration functions
	migrateFunc1 := func() (interface{}, error) {
		// Simulate data migration
		return 1, nil
	}

	migrateFunc2 := func() (interface{}, error) {
		// Simulate data migration
		return 2, nil
	}

	migrateFuncs := []MigrateFunc{migrateFunc1, migrateFunc2}

	// Run the migration with concurrency 2
	concurrency := 2
	futures := Migrate(migrateFuncs, concurrency)

	// Collect and handle results
	var results []int
	for _, future := range futures {
		result, err := future.Get()
		if err != nil {
			// Handle migration errors
			fmt.Println("Migration error:", err)
			continue
		}
		results = append(results, result.(int))
	}

	fmt.Println("Migration completed:", results)
}
