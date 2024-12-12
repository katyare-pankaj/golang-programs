package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ConcurrencyStrategy interface defines the operations that all concurrency strategies must implement.
type ConcurrencyStrategy interface {
	Execute(tasks []func())
}

// GoroutineStrategy implements ConcurrencyStrategy using goroutines.
type GoroutineStrategy struct{}

func (g *GoroutineStrategy) Execute(tasks []func()) {
	for _, task := range tasks {
		go task()
	}
}

// WaitGroupStrategy implements ConcurrencyStrategy using sync.WaitGroup.
type WaitGroupStrategy struct{}

func (w *WaitGroupStrategy) Execute(tasks []func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
		}(task)
	}
	wg.Wait()
}

// ContextStrategy implements ConcurrencyStrategy using context.Context.
type ContextStrategy struct{}

func (c *ContextStrategy) Execute(tasks []func()) {
	ctx := context.Background()
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t func(), ctx context.Context) {
			defer wg.Done()
			t()
		}(task, ctx)
	}

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Context timeout reached.")
	case <-ctx.Done():
		fmt.Println("Context cancelled.")
	}
	wg.Wait()
}

// Task function represents a unit of work to be executed concurrently.
func Task(id int) {
	fmt.Printf("Task %d running...\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Task %d completed.\n", id)
}

func main() {
	tasks := []func(){Task(1), Task(2), Task(3), Task(4), Task(5)}

	fmt.Println("Using GoroutineStrategy:")
	goroutineStrategy := &GoroutineStrategy{}
	goroutineStrategy.Execute(tasks)
	fmt.Println("\nUsing WaitGroupStrategy:")
	waitGroupStrategy := &WaitGroupStrategy{}
	waitGroupStrategy.Execute(tasks)
	fmt.Println("\nUsing ContextStrategy:")
	contextStrategy := &ContextStrategy{}
	contextStrategy.Execute(tasks)
}
