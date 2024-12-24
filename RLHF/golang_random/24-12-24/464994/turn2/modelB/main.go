package main

import (
	"fmt"
)

type Task struct {
	Name string
}

var tasks []Task

func main() {
	showMenu()
}
func showMenu() {
	for {
		fmt.Println(`
        1. Add Task
        2. Remove Task
        3. Display Tasks
        4. Exit
        Enter your choice: `)
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		switch choice {
		case 1:
			addTask()
		case 2:
			removeTask()
		case 3:
			displayTasks()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select between 1-4.")
		}
	}
}
func addTask() {
	var task Task
	fmt.Print("Enter task name: ")
	fmt.Scanln(&task.Name)
	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}
func removeTask() {
	var taskName string
	fmt.Print("Enter task name to remove: ")
	fmt.Scanln(&taskName)
	for index, task := range tasks {
		if task.Name == taskName {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task removed successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}
func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("\nTasks:")
	fmt.Println("--------------------")
	for _, task := range tasks {
		fmt.Printf("%s\n", task.Name)
	}
	fmt.Println("--------------------")
}
