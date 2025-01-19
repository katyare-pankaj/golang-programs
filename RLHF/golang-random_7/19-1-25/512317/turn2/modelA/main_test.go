package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
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
	http.HandleFunc("/increment-mutable", IncrementMutableHandler)
	http.HandleFunc("/increment-immutable", IncrementImmutableHandler)
	http.HandleFunc("/get-count-mutable", GetCountMutableHandler)
	http.HandleFunc("/get-count-immutable", GetCountImmutableHandler)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func IncrementMutableHandler(w http.ResponseWriter, r *http.Request) {
	mutableCounter.Increment()
	fmt.Fprintln(w, "Mutable counter incremented")
}

func IncrementImmutableHandler(w http.ResponseWriter, r *http.Request) {
	immutableCounter = immutableCounter.Increment()
	fmt.Fprintln(w, "Immutable counter incremented")
}

func GetCountMutableHandler(w http.ResponseWriter, r *http.Request) {
	count := mutableCounter.GetCount()
	fmt.Fprintf(w, "Mutable counter: %d\n", count)
}

func GetCountImmutableHandler(w http.ResponseWriter, r *http.Request) {
	count := immutableCounter.GetCount()
	fmt.Fprintf(w, "Immutable counter: %d\n", count)
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
