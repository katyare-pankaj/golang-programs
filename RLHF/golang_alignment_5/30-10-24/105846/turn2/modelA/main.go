package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Counter is a simple shared counter
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// GetCount returns the current count
func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

var counter = &Counter{}

func handler(w http.ResponseWriter, r *http.Request) {
	counter.Increment()
	fmt.Fprintf(w, "Hello, world! Count: %d\n", counter.GetCount())
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
