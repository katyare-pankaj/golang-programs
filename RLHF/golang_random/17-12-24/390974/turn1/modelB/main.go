package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data represents the type of data being processed
type Data string

// Process simulates processing a piece of data
func Process(ctx context.Context, data Data, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Processed: %s\n", data)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Fan-in: Multiple data sources
	dataSources := []chan Data{
		make(chan Data),
		make(chan Data),
		make(chan Data),
	}

	// Workers
	const numWorkers = 5
	workerChan := make(chan Data, numWorkers*10)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				select {
				case data, ok := <-workerChan:
					if !ok {
						return
					}
					wg.Add(1)
					Process(ctx, data, &wg)
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// Simulate data production
	go func() {
		for i := 0; i < 100; i++ {
			for _, source := range dataSources {
				select {
				case source <- Data(fmt.Sprintf("Data %d", i)):
				case <-ctx.Done():
					return
				}
			}
		}
		close(workerChan)
	}()

	// Fan-out: Multiple data sinks
	dataSinks := []chan Data{
		make(chan Data),
		make(chan Data),
	}

	// Distribute processed data to sinks
	go func() {
		for {
			select {
			case data, ok := <-workerChan:
				if !ok {
					return
				}
				for _, sink := range dataSinks {
					select {
					case sink <- data:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All data processed.")
}
