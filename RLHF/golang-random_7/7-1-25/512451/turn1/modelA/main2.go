package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func dependentTask1() {
	fmt.Println("Running dependentTask1")
	wg.Done()
}

func dependentTask2() {
	fmt.Println("Running dependentTask2")
	wg.Done()
}

func dependentTask3() {
	fmt.Println("Running dependentTask3")
	wg.Add(2) // Wait for dependentTask1 and dependentTask2 to complete
	go dependentTask1()
	go dependentTask2()
	wg.Wait() // Wait for the dependencies to finish
	fmt.Println("Running dependentTask3 after dependencies")
	wg.Done()
}

func main() {
	// Function to execute a list of tasks
	executeTasks := func(tasks []func()) {
		for _, task := range tasks {
			go func(t func()) {
				defer wg.Done()
				t()
			}(task)
		}
	}

	// List of dependent tasks
	tasks := []func(){dependentTask3}

	wg.Add(len(tasks))
	executeTasks(tasks)

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All dependent tasks completed.")
}
