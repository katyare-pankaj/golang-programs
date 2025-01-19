package main

import (
	"bytes"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Data structure to test with
type Data struct {
	Value int
}

// Mutable callback example
func mutableCallback(data *Data, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	data.Value = rand.Int()
}

// Immutable callback example
func immutableCallback(data Data, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	_ = Data{Value: rand.Int()}
}

// Test suite
func TestMutableVsImmutable(t *testing.T) {
	const numGoroutines = 1000

	t.Run("TestMemoryUsage", func(t *testing.T) {
		var memStats runtime.MemStats

		// Measure memory usage for mutable state
		runtime.ReadMemStats(&memStats)
		mutableStartAlloc := memStats.Alloc

		var wg sync.WaitGroup
		data := &Data{}
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go mutableCallback(data, &wg)
		}
		wg.Wait()

		runtime.ReadMemStats(&memStats)
		mutableEndAlloc := memStats.Alloc
		mutableMemoryUsage := mutableEndAlloc - mutableStartAlloc
		t.Logf("Mutable Memory Usage: %d bytes", mutableMemoryUsage)

		// Measure memory usage for immutable state
		runtime.ReadMemStats(&memStats)
		immutableStartAlloc := memStats.Alloc

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go immutableCallback(*data, &wg)
		}
		wg.Wait()

		runtime.ReadMemStats(&memStats)
		immutableEndAlloc := memStats.Alloc
		immutableMemoryUsage := immutableEndAlloc - immutableStartAlloc
		t.Logf("Immutable Memory Usage: %d bytes", immutableMemoryUsage)
	})

	t.Run("TestConcurrency", func(t *testing.T) {
		// Mutable state concurrency test
		var wg sync.WaitGroup
		data := &Data{}
		mutableResults := make(chan int, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				mutableCallback(data, &wg)
				mutableResults <- data.Value
			}()
		}
		wg.Wait()
		close(mutableResults)

		mutableValueSet := make(map[int]bool)
		for v := range mutableResults {
			mutableValueSet[v] = true
		}
		t.Logf("Mutable unique values: %d", len(mutableValueSet))

		// Immutable state concurrency test
		immutableResults := make(chan int, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(d Data) {
				immutableCallback(d, &wg)
				immutableResults <- d.Value
			}(*data)
		}
		wg.Wait()
		close(immutableResults)

		immutableValueSet := make(map[int]bool)
		for v := range immutableResults {
			immutableValueSet[v] = true
		}
		t.Logf("Immutable unique values: %d", len(immutableValueSet))
	})

	t.Run("TestDataIntegrity", func(t *testing.T) {
		// Testing data integrity with mutable state
		var wg sync.WaitGroup
		data := &Data{}
		mutableBuffer := bytes.Buffer{}

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				mutableCallback(data, &wg)
				mutableBuffer.WriteString(string(data.Value))
			}()
		}
		wg.Wait()

		// Testing data integrity with immutable state
		immutableBuffer := bytes.Buffer{}

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(d Data) {
				immutableCallback(d, &wg)
				immutableBuffer.WriteString(string(d.Value))
			}(*data)
		}
		wg.Wait()

		if mutableBuffer.String() != immutableBuffer.String() {
			t.Log("Data integrity test passed: Buffers are different")
		} else {
			t.Error("Data integrity test failed: Buffers are the same")
		}
	})
}
