package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	lock sync.Mutex
	name string
}

func (r *Resource) Acquire() {
	r.lock.Lock()
}

func (r *Resource) Release() {
	r.lock.Unlock()
}

func useResource(wg *sync.WaitGroup, r *Resource, name string) {
	defer wg.Done()

	// Acquire the resource
	r.Acquire()
	defer r.Release()

	// Simulate using the resource
	fmt.Printf("Goroutine %s using resource %s\n", name, r.name)
	time.Sleep(2 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	r := &Resource{name: "Shared Resource"}

	// Create multiple Goroutines that use the shared resource
	wg.Add(3)
	go useResource(&wg, r, "Goroutine 1")
	go useResource(&wg, r, "Goroutine 2")
	go useResource(&wg, r, "Goroutine 3")

	// Wait for all Goroutines to complete
	wg.Wait()

	fmt.Println("All Goroutines have completed.")
}
