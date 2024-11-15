// Golang example to demonstrate parallel data access and processing with portable memory management.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	dataSize   = 1000000 // Simulated data size
	numWorkers = 4       // Number of parallel workers
)

// SimulatedData represents the data we want to process.
type SimulatedData struct {
	values []float64
}

// CreateData generates the value to be processed
func (s *SimulatedData) createData() {
	s.values = make([]float64, dataSize)
	for i := range s.values {
		s.values[i] = rand.Float64()
	}
}

// ProcessDataProcess a chunk of data
func processData(data []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	// Worker processing logic goes here
}

func main() {
	rand.Seed(time.Now().UnixNano())

	data := &SimulatedData{}
	data.createData()
	fmt.Println("Data creation completed")

	// Break the data into chunks for parallel processing
	chunkSize := dataSize / numWorkers
	dataChunks := make([][]float64, numWorkers)
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, dataSize)
		dataChunks[i] = data.values[start:end]
	}

	var wg sync.WaitGroup

	// Start parallel workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processData(dataChunks[i], &wg)
	}

	wg.Wait()
	fmt.Println("Data processing completed")

	// Memory cleanup (automatic in Go due to garbage collection, but good practice to nil slices)
	data.values = nil
	for i := range dataChunks {
		dataChunks[i] = nil
	}

	fmt.Println("Memory cleaned up")
	runtime.GC() // Explicit GC call is optional but can help ensure immediate cleanup
	fmt.Println("GC completed")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
