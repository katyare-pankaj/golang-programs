package main

import (
	"fmt"
	"sort"
	"sync"
)

// DataPoint represents a point in the dataset with a float feature and a label.
type DataPoint struct {
	Feature float64
	Label   string
}

// parallelSort sorts the dataset in chunks and merges them.
func parallelSort(data []DataPoint) {
	numGoroutines := 4
	chunkSize := (len(data) + numGoroutines - 1) / numGoroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		go func(dataChunk []DataPoint) {
			defer wg.Done()
			sort.Slice(dataChunk, func(i, j int) bool {
				return dataChunk[i].Feature < dataChunk[j].Feature
			})
		}(data[start:end])
	}

	wg.Wait()

	// Merge sorted chunks (this step can be optimized further)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Feature < data[j].Feature
	})
}

func main() {
	// Sample large dataset
	data := []DataPoint{
		{Feature: 2.3, Label: "B"},
		{Feature: 1.1, Label: "A"},
		{Feature: 4.5, Label: "D"},
		{Feature: 3.0, Label: "C"},
		// Add more data points for a larger dataset
	}

	// Parallel sort
	parallelSort(data)
	fmt.Println("Sorted data:", data)
}
