package main

import (
	"fmt"
	"os"
)

type ToDoList []string

func (tdl *ToDoList) AddTask(task string) {
	*tdl = append(*tdl, task)
}

func (tdl *ToDoList) RemoveTask(index int) {
	if index >= 0 && index < len(*tdl) {
		*tdl = append((*tdl)[:index], (*tdl)[index+1:]...)
	}
}

func (tdl ToDoList) ListTasks() {
	for index, task := range tdl {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func main() {
	var toDoList ToDoList

	for {
		fmt.Println("To-Do List Manager:")
		fmt.Println("-------------------")
		fmt.Println("1. Add a task")
		fmt.Println("2. Remove a task")
		fmt.Println("3. List tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter a task to add: ")
			var task string
			_, err := fmt.Scanln(&task)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task.")
				continue
			}
			toDoList.AddTask(task)
		case 2:
			fmt.Print("Enter the index of the task to remove: ")
			var index int
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid index.")
				continue
			}
			toDoList.RemoveTask(index - 1)
		case 3:
			fmt.Println("Your To-Do List:")
			toDoList.ListTasks()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
