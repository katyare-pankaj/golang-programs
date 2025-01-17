package main

import (
	"math/rand"
	"testing"
	"time"
)

// ProcessData simulates processing a batch of data. Each operation represents a real-world task.
func ProcessData(data []int) []int {
	for i := range data {
		// Simulate a task by sleeping for a random short duration.
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		// An example operation: increment each element.
		data[i]++
	}
	return data
}

func generateTestData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	return data
}

// BenchmarkProcessData benchmarks the ProcessData function.
func BenchmarkProcessData(b *testing.B) {
	data := generateTestData(1000) // Adjust size to match real-world batch sizes.
	for n := 0; n < b.N; n++ {
		ProcessData(data)
	}
}

func MeasureUserExperience(data []int) {
	start := time.Now()
	ProcessData(data)
	elapsed := time.Since(start)
	// In real-world scenarios, log or handle the elapsed time appropriately.
	println("User Experience Duration:", elapsed.Milliseconds(), "ms")
}

func main() {
	// Run the user experience measurement with test data.
	testData := generateTestData(1000)
	MeasureUserExperience(testData)
}
