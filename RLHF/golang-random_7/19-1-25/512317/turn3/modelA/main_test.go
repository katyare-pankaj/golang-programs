package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"

	"github.com/labstack/echo/v4"
)

// MutableCounter with a mutex for concurrency
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

// ImmutableCounter returning a new instance upon increment
type ImmutableCounter struct {
	count int
	mu    sync.RWMutex
}

func (c *ImmutableCounter) Increment() *ImmutableCounter {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return &ImmutableCounter{count: c.count + 1}
}

func (c *ImmutableCounter) GetCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

var (
	mutableCounter   = &MutableCounter{}
	immutableCounter = &ImmutableCounter{}
)

func main() {
	e := echo.New()

	e.GET("/increment-mutable", func(c echo.Context) error {
		mutableCounter.Increment()
		return c.String(http.StatusOK, "Mutable counter incremented")
	})

	e.GET("/increment-immutable", func(c echo.Context) error {
		immutableCounter = immutableCounter.Increment()
		return c.String(http.StatusOK, "Immutable counter incremented")
	})

	e.GET("/get-count-mutable", func(c echo.Context) error {
		count := mutableCounter.GetCount()
		return c.String(http.StatusOK, fmt.Sprintf("Mutable counter: %d", count))
	})

	e.GET("/get-count-immutable", func(c echo.Context) error {
		count := immutableCounter.GetCount()
		return c.String(http.StatusOK, fmt.Sprintf("Immutable counter: %d", count))
	})

	log.Println("Starting server at :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func TestConcurrencyMutable(t *testing.T) {
	var wg sync.WaitGroup
	testRequests := 10000

	for i := 0; i < testRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutableCounter.Increment()
		}()
	}

	wg.Wait()
	if getVal := mutableCounter.GetCount(); getVal != testRequests {
		t.Errorf("Expected %d, got %d", testRequests, getVal)
	}
}

func TestConcurrencyImmutable(t *testing.T) {
	var wg sync.WaitGroup
	testRequests := 10000

	for i := 0; i < testRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			immutableCounter = immutableCounter.Increment()
		}()
	}

	wg.Wait()
	if getVal := immutableCounter.GetCount(); getVal != testRequests {
		t.Errorf("Expected %d, got %d", testRequests, getVal)
	}
}

func BenchmarkMutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutableCounter.Increment()
	}
}

func BenchmarkImmutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		immutableCounter = immutableCounter.Increment()
	}
}
