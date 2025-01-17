package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// ProcessTask simulates processing an individual task.
func ProcessTask(data int) int {
	// Simulate a delay to mimic a real-world task.
	delay := time.Millisecond * time.Duration(rand.Intn(10))
	time.Sleep(delay)
	return data + 1
}

// ProcessBatch processes a batch of data, measuring latency for each individual task.
func ProcessBatch(data []int) []int {
	var wg sync.WaitGroup
	output := make([]int, len(data))
	latencies := make([]time.Duration, len(data))

	for i := range data {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			output[i] = ProcessTask(data[i])
			latencies[i] = time.Since(start)
		}(i)
	}

	wg.Wait()

	// Calculate average latency and throughput
	var totalLatency time.Duration
	for _, latency := range latencies {
		totalLatency += latency
	}
	avgLatency := totalLatency / time.Duration(len(data))
	throughput := float64(len(data)) / totalLatency.Seconds()

	fmt.Printf("Average Latency: %v\n", avgLatency)
	fmt.Printf("Throughput: %.2f tasks/second\n", throughput)

	return output
}

func generateLargeTestData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

// BenchmarkProcessBatch benchmarks the ProcessBatch function.
func BenchmarkProcessBatch(b *testing.B) {
	data := generateLargeTestData(10000) // Large batch size
	for n := 0; n < b.N; n++ {
		ProcessBatch(data)
	}
}

func main() {
	// Manually test to observe results, outside the benchmark context
	testData := generateLargeTestData(10000)
	ProcessBatch(testData)
}
