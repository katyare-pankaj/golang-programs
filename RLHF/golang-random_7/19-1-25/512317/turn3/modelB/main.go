package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// MutableCounter using an atomic int64 for concurrency-safe updates
type MutableCounter struct {
	count int64
}

func (c *MutableCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *MutableCounter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

// ImmutableCounter using an atomic pointer for concurrency-safe updates
type ImmutableCounter struct {
	count *int
}

func (c *ImmutableCounter) Increment() *ImmutableCounter {
	newCount := *c.count + 1
	return &ImmutableCounter{&newCount}
}

func (c *ImmutableCounter) GetCount() int {
	return *c.count
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize mutable and immutable counters
	mutableCounter := &MutableCounter{}
	immutableCounter := &ImmutableCounter{count: new(int)}

	// Mutable counter endpoint
	e.GET("/mutable-increment", func(c echo.Context) error {
		mutableCounter.Increment()
		return c.String(http.StatusOK, "Mutable counter incremented")
	})

	e.GET("/mutable-count", func(c echo.Context) error {
		count := mutableCounter.GetCount()
		return c.String(http.StatusOK, fmt.Sprintf("Mutable counter: %d", count))
	})

	// Immutable counter endpoint
	e.GET("/immutable-increment", func(c echo.Context) error {
		immutableCounter = immutableCounter.Increment()
		return c.String(http.StatusOK, "Immutable counter incremented")
	})

	e.GET("/immutable-count", func(c echo.Context) error {
		count := immutableCounter.GetCount()
		return c.String(http.StatusOK, fmt.Sprintf("Immutable counter: %d", count))
	})

	// Start the server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// benchmarking
	benchmarkResults := make(chan time.Duration)
	go benchmarkMutable(mutableCounter, benchmarkResults)
	go benchmarkImmutable(immutableCounter, benchmarkResults)

	for i := 0; i < 2; i++ {
		duration := <-benchmarkResults
		fmt.Printf("Benchmark Result %d: %s\n", i+1, duration)
	}
}

func benchmarkMutable(counter *MutableCounter, result chan<- time.Duration) {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		counter.Increment()
	}
	result <- time.Since(start)
}

func benchmarkImmutable(counter *ImmutableCounter, result chan<- time.Duration) {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		counter = counter.Increment()
	}
	result <- time.Since(start)
}
