package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task struct represents an item to be picked and packed
type Task struct {
	ID       int
	Name     string
	Weight   int // in kg
	Picking  bool
	Packing  bool
	Duration time.Duration // time taken to process
}

// SimulatePicking simulates picking an item from the warehouse
func SimulatePicking(task *Task) {
	time.Sleep(task.Duration)
	task.Picking = true
	fmt.Printf("Picked item %s (weight: %d kg)\n", task.Name, task.Weight)
}

// SimulatePacking simulates packing an item
func SimulatePacking(task *Task) {
	time.Sleep(task.Duration)
	task.Packing = true
	fmt.Printf("Packed item %s (weight: %d kg)\n", task.Name, task.Weight)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize tasks
	tasks := []Task{}
	for i := 1; i <= 10; i++ {
		task := Task{
			ID:       i,
			Name:     fmt.Sprintf("Item%d", i),
			Weight:   rand.Intn(10) + 1,
			Picking:  false,
			Packing:  false,
			Duration: time.Duration(rand.Intn(10)+1) * time.Second,
		}
		tasks = append(tasks, task)
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Start picking and packing goroutines
	for _, task := range tasks {
		wg.Add(1) // Increment the WaitGroup count
		go func(t Task) {
			defer wg.Done() // Decrement the WaitGroup count when the goroutine is complete
			// Simulate picking
			SimulatePicking(&t)

			// Simulate packing
			SimulatePacking(&t)
		}(task)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check the status of all tasks
	fmt.Println("\nAll items have been picked and packed.")
}
