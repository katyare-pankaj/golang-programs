package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Define a result structure with result and error interface
type Result struct {
	Result interface{}
	Error  error
}

// Task type as a function that returns a Result
type Task func() *Result

// Function to run tasks concurrently using Goroutines and collect results
func RunTasksConcurrently(tasks []Task) []Result {
	var results []Result
	var wg sync.WaitGroup
	resultChan := make(chan *Result)

	// Worker function to run a task and send its result to a channel
	runTask := func(task Task) {
		defer wg.Done()
		results <- task()
	}

	// Start Goroutines for each task
	for _, task := range tasks {
		wg.Add(1)
		go runTask(task)
	}

	// Collect results from the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Read results from the channel
	for result := range resultChan {
		results = append(results, *result)
	}

	return results
}

// Define example tasks
func task1() *Result {
	time.Sleep(1 * time.Second)
	return &Result{Result: "task1 successful", Error: nil}
}

func task2() *Result {
	time.Sleep(1 * time.Second)
	return &Result{Result: nil, Error: errors.New("task2 failed")}
}

func task3() *Result {
	time.Sleep(2 * time.Second)
	return &Result{Result: "task3 successful", Error: nil}
}

func task4() *Result {
	time.Sleep(1 * time.Second)
	return &Result{Result: nil, Error: errors.New("task4 network error")}
}

func main() {
	// Define tasks to run
	tasks := []Task{task1, task2, task3, task4}

	// Run tasks concurrently and collect results
	results := RunTasksConcurrently(tasks)

	// Process results in a thread-safe manner
	for _, result := range results {
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
		} else {
			fmt.Println("Result:", result.Result)
		}
	}
}
