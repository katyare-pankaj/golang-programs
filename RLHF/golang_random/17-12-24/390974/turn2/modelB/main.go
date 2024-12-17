package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data represents the data being streamed
type Data struct {
	ID    int
	Value float64
}

// Process simulates processing a piece of data
func Process(ctx context.Context, data Data, wg *sync.WaitGroup, results chan<- Data) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		// Simulate processing time
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		// Transform or process data
		processedData := Data{ID: data.ID, Value: data.Value * 2.0}
		results <- processedData
		fmt.Printf("Processed: %+v\n", processedData)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Channel for incoming data stream
	dataIn := make(chan Data, 100)

	// Channel for processed data
	dataOut := make(chan Data)

	// Worker pool
	const numWorkers = 4
	workerPool := make(chan struct{}, numWorkers)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				select {
				case <-workerPool:
					// Defer returning the worker to the pool
					defer func() { workerPool <- struct{}{} }()

					select {
					case data, ok := <-dataIn:
						if !ok {
							return
						}
						wg.Add(1)
						Process(ctx, data, &wg, dataOut)
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// Simulate data streaming
	go func() {
		for i := 0; i < 1000; i++ {
			dataIn <- Data{ID: i, Value: float64(i) * 10.0}
			time.Sleep(10 * time.Millisecond) // Simulate incoming data rate
		}
		close(dataIn)
	}()

	// Consume processed data
	go func() {
		for {
			select {
			case processedData, ok := <-dataOut:
				if !ok {
					return
				}
				fmt.Printf("Consumed: %+v\n", processedData)
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All data processed.")
}
