package main

import (
	"fmt"
	"sort"
)

// Task represents a simple task with a title and a due date.
type Task struct {
	Title   string
	DueDate string
}

// TaskList is a slice of tasks that represents the list of tasks.
type TaskList []Task

// AddTask adds a new task to the task list.
func (t *TaskList) AddTask(title, dueDate string) {
	*t = append(*t, Task{Title: title, DueDate: dueDate})
}

// SortTasks sorts the task list based on their due date.
func (t *TaskList) SortTasks() {
	sort.Slice(*t, func(i, j int) bool {
		return (*t)[i].DueDate < (*t)[j].DueDate
	})
}

// DisplayTasks prints out the task list in a organized manner.
func (t *TaskList) DisplayTasks() {
	for _, task := range *t {
		fmt.Printf("Task: %s, Due Date: %s\n", task.Title, task.DueDate)
	}
}

func main() {
	var taskList TaskList
	taskList.AddTask("Learn Go", "2023-07-30")
	taskList.AddTask("Implement Task Manager", "2023-08-10")
	taskList.AddTask("Buy groceries", "2023-07-25")

	fmt.Println("--- Unsorted Tasks ---")
	taskList.DisplayTasks()

	fmt.Println("--- Sorted Tasks ---")
	taskList.SortTasks()
	taskList.DisplayTasks()
}
