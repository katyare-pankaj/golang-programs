package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value int
}

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan Data, 10)
	transformedChan := make(chan Data, 10)
	aggregatedChan := make(chan Data, 10)

	// Producer goroutines to generate Data and send it to dataChan
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				data := Data{Value: i*10 + j}
				dataChan <- data
				time.Sleep(time.Millisecond * 100) // Simulate some work
			}
		}()
	}

	// Transformer goroutines to receive Data from dataChan, transform it, and send it to transformedChan
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range dataChan {
				data.Value += 100 // Transform the data by incrementing
				transformedChan <- data
			}
		}()
	}

	// Aggregator goroutine to receive transformed data from transformedChan, aggregate it, and print the results
	go func() {
		var sum int
		var count int

		for data := range transformedChan {
			sum += data.Value
			count++
		}

		fmt.Printf("Aggregated Results: Sum = %d, Count = %d\n", sum, count)
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(aggregatedChan)
}
