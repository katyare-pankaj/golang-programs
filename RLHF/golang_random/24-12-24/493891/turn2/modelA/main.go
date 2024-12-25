package main

import (
	"fmt"
)

// Task struct represents a task with a description and a completion status.
type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func main() {
	showMenu()
}

// showMenu presents the main menu options to the user.
func showMenu() {
	for {
		fmt.Println(`
        1. Add Task
        2. Remove Task
        3. Display Tasks
        4. Complete Task
        5. Exit
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
			completeTask()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select between 1-5.")
		}
	}
}

// addTask adds a new task to the tasks slice.
func addTask() {
	var task Task
	fmt.Print("Task Description: ")
	fmt.Scanln(&task.Description)
	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}

// removeTask removes a task from the tasks slice by its index.
func removeTask() {
	var index int
	displayTasks()
	fmt.Print("Enter the index of the task to remove: ")
	_, err := fmt.Scanf("%d", &index)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid index. Please try again.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task removed successfully.")
}

// displayTasks lists all the tasks in the tasks slice with their details.
func displayTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	fmt.Println("\nTasks:")
	fmt.Println("--------------------")
	for index, task := range tasks {
		status := "[Incomplete]"
		if task.Completed {
			status = "[Completed]"
		}
		fmt.Printf("%d. %s - %s\n", index, task.Description, status)
	}
	fmt.Println("--------------------")
}

// completeTask marks a task as completed by its index.
func completeTask() {
	var index int
	displayTasks()
	fmt.Print("Enter the index of the task to complete: ")
	_, err := fmt.Scanf("%d", &index)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid index. Please try again.")
		return
	}
	tasks[index-1].Completed = !tasks[index-1].Completed
	fmt.Println("Task completed successfully.")
}
