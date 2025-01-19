
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

// Mutable state implementation
type MutableCounter struct {
	count int
	mu    sync.Mutex
}

func (c *MutableCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutableCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// Immutable state implementation
type ImmutableCounter struct {
	count int
}

func (c *ImmutableCounter) Increment() *ImmutableCounter {
	return &ImmutableCounter{count: c.count + 1}
}

func (c *ImmutableCounter) GetCount() int {
	return c.count
}

var (
	mutableCounter  = &MutableCounter{}
	immutableCounter = &ImmutableCounter{}
)

// Echo web application handler for mutable counter
func mutableHandler(c echo.Context) error {
	mutableCounter.Increment()
	return c.String(http.StatusOK, fmt.Sprintf("Mutable Counter: %d", mutableCounter.GetCount()))
}

// Echo web application handler for immutable counter
func immutableHandler(c echo.Context) error {
	immutableCounter = immutableCounter.Increment()
	return c.String(http.StatusOK, fmt.Sprintf("Immutable Counter: %d", immutableCounter.GetCount()))
}

func main() {
	e := echo.New()

	// Routes for mutable and immutable counters
	e.GET("/mutable", mutableHandler)
	e.GET("/immutable", immutableHandler)

	// Start the web server
	go e.Logger.Fatal(e.Start(":8080"))

	// Perform benchmarking and concurrency testing
	benchmarkAndTest()
}

func benchmarkAndTest() {
	// Benchmarking
	fmt.Println("Benchmarking...")
	b.Run("Mutable", BenchmarkMutableCallbacks)
	b.Run("Immutable", BenchmarkImmutableCallbacks)

	// Concurrency testing
	fmt.Println("\nConcurrency Testing...")
	testCallbackIntegrity()
}

func BenchmarkMutableCallbacks(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutableCounter.Increment()
	}
}

func BenchmarkImmutableCallbacks(b *testing.B) {
	for n := 0; n < b.N; n++ {
		immutableCounter = immutableCounter.Increment()
	}
}

func testCallbackIntegrity() {
	// Perform concurrent increment operations on both counters
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for n := 0; n < 10000; n++ {
			mutableCounter.Increment()
		}
	}()

	go func() {
		defer wg.Done()
		for n := 0; n < 10000; n++ {
			mutableCounter.Increment()
		}
	}()

	wg.Wait()
	expectedCount := 20000
	actualCount := mutableCounter.GetCount()
	if actualCount != expectedCount {
		fmt.Println("Mutable counter integrity failed: Expected", expectedCount, ", got", actualCount)
	} else {
		fmt.Println("Mutable counter integrity passed.")
	}

	// Test immutable counter