package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Input  int
	Result int
}

type ResultAggregator struct {
	mu      sync.Mutex
	results []int
}

func (a *ResultAggregator) AddResult(result int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.results = append(a.results, result)
}

func (a *ResultAggregator) GetResults() []int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.results
}

func worker(wg *sync.WaitGroup, tasks chan Task, aggregator *ResultAggregator) {
	defer wg.Done()
	for task := range tasks {
		// Simulate some work with a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		result := task.Input * task.Input
		aggregator.AddResult(result)
		fmt.Printf("Task %d completed: %d * %d = %d\n", task.Input, task.Input, task.Input, result)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numTasks := 20
	tasks := make(chan Task, numTasks)

	aggregator := &ResultAggregator{}

	// Start workers
	var wg sync.WaitGroup
	numWorkers := 5 // Adjust the number of workers as needed
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, aggregator)
	}

	// Send tasks to the worker pool
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{Input: i}
	}
	close(tasks)

	// Wait for all tasks to complete
	wg.Wait()

	// Summarize results
	fmt.Println("\nAll tasks completed. Aggregating results:")
	results := aggregator.GetResults()
	var totalResult int
	for _, result := range results {
		totalResult += result
	}
	fmt.Printf("\nTotal sum of results: %d\n", totalResult)
}
