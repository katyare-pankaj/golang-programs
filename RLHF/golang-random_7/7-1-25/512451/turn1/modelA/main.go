package main

import (
	"fmt"
	"sync"
)

func independentTask1() {
	fmt.Println("Running independentTask1")
}

func independentTask2() {
	fmt.Println("Running independentTask2")
}

func independentTask3() {
	fmt.Println("Running independentTask3")
}

func main() {
	var wg sync.WaitGroup

	// Function to execute a list of tasks
	executeTasks := func(tasks []func()) {
		for _, task := range tasks {
			wg.Add(1)
			go func(t func()) {
				defer wg.Done()
				t()
			}(task)
		}
	}

	// List of independent tasks
	tasks := []func(){independentTask1, independentTask2, independentTask3}

	executeTasks(tasks)

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All independent tasks completed.")
}
