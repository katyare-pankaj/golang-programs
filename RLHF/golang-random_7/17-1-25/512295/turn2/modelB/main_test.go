package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	numIterations = 1000 // Number of iterations for each batch
	batchSize     = 100  // Size of each batch
)

func benchmarkProcessData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataBatch := generateRandomDataBatch()
		processData(dataBatch)
	}
}

func measureUserExperience(b *testing.B) {
	b.StopTimer()
	defer b.StartTimer()

	var waitGroup sync.WaitGroup
	waitGroup.Add(numIterations)

	start := time.Now()
	for i := 0; i < numIterations; i++ {
		go func() {
			dataBatch := generateRandomDataBatch()
			processData(dataBatch)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	end := time.Now()
	fmt.Printf("User Experience Duration: %v\n", end.Sub(start))
}

// Generates a random batch of data for testing purposes
func generateRandomDataBatch() []int {
	batch := make([]int, batchSize)
	for i := range batch {
		batch[i] = rand.Intn(1000)
	}
	return batch
}

// Simulates real-world data processing tasks
func processData(batch []int) {
	for _, data := range batch {
		// Simulate a task that involves data processing and loading a remote resource
		time.Sleep(time.Duration(data) * time.Millisecond)
		_ = data + 10 // Do some basic arithmetic
	}
}

func main() {
	testing.Benchmark(benchmarkProcessData)
	measureUserExperience(&testing.B{})
}
