package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	batchSize       = 1000
	numBatches      = 1000
	taskDelayMillis = 5
)

func ProcessData(data []int) []int {
	for i := range data {
		// Simulate a task by sleeping for a random short duration.
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(taskDelayMillis)))
		// An example operation: increment each element.
		data[i]++
	}
	return data
}

func generateTestData() []int {
	data := make([]int, batchSize)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

// BenchmarkProcessData benchmarks the ProcessData function.
func BenchmarkProcessData(b *testing.B) {
	data := generateTestData()
	b.SetBytes(int64(len(data)) * int64(numBatches))

	for n := 0; n < b.N; n++ {
		// Measure the time taken to process one batch.
		start := time.Now()
		ProcessData(data)
		elapsed := time.Since(start)

		// Report the latency for each task in milliseconds.
		latency := float64(elapsed.Nanoseconds()) / float64(batchSize) / 1000000
		b.ReportMetric(latency, "latency-ms")
	}
}

func main() {
}
