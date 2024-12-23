package main

import (
	"fmt"
)

// Task represents a task with a name and priority
type Task struct {
	Name     string
	Priority int
}

// TaskManager manages a list of tasks
type TaskManager struct {
	Tasks []Task
}

// AddTask adds a new task to the manager
func (t *TaskManager) AddTask(name string, priority int) {
	t.Tasks = append(t.Tasks, Task{Name: name, Priority: priority})
}

// DisplayTasks prints all tasks in the manager
func (t *TaskManager) DisplayTasks() {
	for _, task := range t.Tasks {
		fmt.Printf("Task: %s (Priority: %d)\n", task.Name, task.Priority)
	}
}

// MarkTaskComplete removes a task from the list by index
func (t *TaskManager) MarkTaskComplete(index int) {
	if index >= 0 && index < len(t.Tasks) {
		t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
		fmt.Printf("Task marked as complete: %s\n", t.Tasks[index].Name)
	} else {
		fmt.Println("Invalid task index.")
	}
}

// Main function to demonstrate the task manager
func main() {
	taskManager := TaskManager{}

	// Add a few tasks
	taskManager.AddTask("Buy groceries", 1)
	taskManager.AddTask("Read book", 2)
	taskManager.AddTask("Go for a walk", 3)

	// Display all tasks
	fmt.Println("Current tasks:")
	taskManager.DisplayTasks()

	// Mark a task as complete
	fmt.Println("\nMarking a task as complete...")
	taskManager.MarkTaskComplete(1)

	// Display tasks again
	fmt.Println("\nUpdated tasks:")
	taskManager.DisplayTasks()
}
