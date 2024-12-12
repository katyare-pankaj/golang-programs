package main

import (
	"fmt"
	"sync"
)

// Task represents a single unit of work.
type Task func(int)

// Strategy defines the interface for different concurrency strategies.
type Strategy interface {
	Execute(tasks []Task)
}

// SequentialStrategy executes tasks sequentially.
type SequentialStrategy struct{}

func (s *SequentialStrategy) Execute(tasks []Task) {
	for i, task := range tasks {
		fmt.Printf("Executing task %d sequentially.\n", i)
		task(i)
	}
}

// ParallelStrategy executes tasks concurrently.
type ParallelStrategy struct{}

func (p *ParallelStrategy) Execute(tasks []Task) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("Executing task %d concurrently.\n", idx)
			task(idx)
		}(i)
	}
	wg.Wait()
}

// Context is a context to use a concurrency strategy.
type Context struct {
	strategy Strategy
}

// SetStrategy sets the strategy for the context.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteTasks runs the selected concurrency strategy.
func (c *Context) ExecuteTasks(tasks []Task) {
	c.strategy.Execute(tasks)
}

// Example Task
func main() {
	tasks := []Task{
		func(i int) { fmt.Println("Task:", i, "completed.") },
		func(i int) { fmt.Println("Task:", i, "completed.") },
		func(i int) { fmt.Println("Task:", i, "completed.") },
	}

	// Create a context
	context := &Context{}

	// Set the strategy to Sequential and execute
	seqStrategy := &SequentialStrategy{}
	context.SetStrategy(seqStrategy)
	fmt.Println("Using Sequential Strategy:")
	context.ExecuteTasks(tasks)

	// Set the strategy to Parallel and execute
	parStrategy := &ParallelStrategy{}
	context.SetStrategy(parStrategy)
	fmt.Println("\nUsing Parallel Strategy:")
	context.ExecuteTasks(tasks)
}
