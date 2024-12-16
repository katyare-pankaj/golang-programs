package main

import (
	"fmt"
	"math"
	"sync"
)

// Define a data chunk
type DataChunk struct {
	data []float64
	mean float64
	std  float64
}

// Worker function processes a data chunk
func processChunk(ch <-chan DataChunk, results chan<- DataChunk, wg *sync.WaitGroup) {
	defer wg.Done()
	for chunk := range ch {
		n := float64(len(chunk.data))
		if n == 0 {
			continue
		}

		sum := 0.0
		sumSq := 0.0
		for _, value := range chunk.data {
			sum += value
			sumSq += value * value
		}

		chunk.mean = sum / n
		chunk.std = math.Sqrt((sumSq - sum*sum/n) / n)

		results <- chunk
	}
}

func main() {
	const numWorkers = 8 // Number of worker goroutines

	// Large dataset
	dataset := make([]float64, 100000000)
	for i := range dataset {
		dataset[i] = float64(i)
	}

	// Divide the dataset into chunks
	chunkSize := 100000
	var chunks []DataChunk
	for i := 0; i < len(dataset); i += chunkSize {
		end := min(i+chunkSize, len(dataset))
		chunks = append(chunks, DataChunk{data: dataset[i:end]})
	}

	// Create input and output channels
	ch := make(chan DataChunk, numWorkers*2)
	results := make(chan DataChunk, numWorkers*2)

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processChunk(ch, results, &wg)
	}

	// Feed chunks into the worker channel
	go func() {
		for _, chunk := range chunks {
			ch <- chunk
		}
		close(ch)
	}()

	// Aggregate results
	var totalMean, totalSumSq float64
	var totalN int
	for result := range results {
		totalMean += result.mean * float64(len(result.data))
		totalSumSq += result.std*result.std*float64(len(result.data)) + result.mean*result.mean*float64(len(result.data))
		totalN += len(result.data)
	}

	overallMean := totalMean / float64(totalN)
	overallStd := math.Sqrt((totalSumSq - totalMean*totalMean/float64(totalN)) / float64(totalN))

	wg.Wait()
	close(results)

	fmt.Printf("Overall Mean: %.2f, Overall Std: %.2f\n", overallMean, overallStd)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
