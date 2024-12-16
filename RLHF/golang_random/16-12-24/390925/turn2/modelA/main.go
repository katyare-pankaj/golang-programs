package main

import (
	"fmt"
	"reflect"
	"sync"
)

// Task is an interface that all tasks must implement
type Task interface {
	Execute()
}

// PrintTask simply prints a message
type PrintTask struct {
	Message string
}

func (p *PrintTask) Execute() {
	fmt.Println("Executing PrintTask:", p.Message)
}

// MultiplyTask multiplies two numbers
type MultiplyTask struct {
	A, B int
}

func (m *MultiplyTask) Execute() {
	result := m.A * m.B
	fmt.Printf("Executing MultiplyTask: %d * %d = %d\n", m.A, m.B, result)
}

// DynamicTaskHandler processes tasks of various types using reflection
func DynamicTaskHandler(tasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		valueOfTask := reflect.ValueOf(task)

		// Call the Execute method using reflection.
		executeMethod := valueOfTask.MethodByName("Execute")
		if executeMethod.IsValid() {
			executeMethod.Call(nil)
		} else {
			fmt.Println("Error: No valid Execute method found.")
		}
	}
}

func main() {
	tasks := make(chan Task)

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the dynamic task handler in a goroutine
	go DynamicTaskHandler(tasks, &wg)

	// Create tasks and push them to the channel
	tasks <- &PrintTask{Message: "Hello, World!"}
	tasks <- &MultiplyTask{A: 10, B: 5}

	// Close the task channel after sending all tasks
	close(tasks)

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("All tasks done!")
}
