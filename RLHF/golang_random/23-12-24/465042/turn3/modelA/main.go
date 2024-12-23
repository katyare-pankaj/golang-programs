package main

import (
	"fmt"
)

// Task represents a single task in the to-do list.
type Task struct {
	Description string
	Completed   bool
}

// ToDoList is a slice of Task that manages the to-do list.
type ToDoList []Task

// AddTask adds a new task to the to-do list.
func (t *ToDoList) AddTask(description string) {
	*t = append(*t, Task{Description: description, Completed: false})
}

// RemoveTask removes a task at a specific index from the to-do list.
func (t *ToDoList) RemoveTask(index int) {
	if index < 0 || index >= len(*t) {
		fmt.Println("Invalid task index.")
		return
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
}

// MarkTaskAsCompleted marks a task at a specific index as completed.
func (t *ToDoList) MarkTaskAsCompleted(index int) {
	if index < 0 || index >= len(*t) {
		fmt.Println("Invalid task index.")
		return
	}
	(*t)[index].Completed = true
}

// ListTasks prints out all tasks in the to-do list.
func (t *ToDoList) ListTasks() {
	for _, task := range *t {
		completedStatus := "❌"
		if task.Completed {
			completedStatus = "✅"
		}
		fmt.Printf("[%s] %s\n", completedStatus, task.Description)
	}
}

func main() {
	var toDoList ToDoList

	// Adding tasks
	toDoList.AddTask("Buy groceries")
	toDoList.AddTask("Learn Go")
	toDoList.AddTask("Read a book")

	fmt.Println("--- Initial To-Do List ---")
	toDoList.ListTasks()

	// Marking a task as completed
	toDoList.MarkTaskAsCompleted(1)

	fmt.Println("--- After Marking Task 2 as Completed ---")
	toDoList.ListTasks()

	// Removing a task
	toDoList.RemoveTask(0)

	fmt.Println("--- After Removing Task 1 ---")
	toDoList.ListTasks()
}
