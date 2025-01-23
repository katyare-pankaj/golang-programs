package main

import (
	"log"
	"net/http"

	"google.golang.org/api/tasks/v1"
)

// ... (same code as before)

func showTasks(w http.ResponseWriter, r *http.Request) {
	// Create a slice of pending tasks
	pendingTasks := []Task{}

	for _, task := range tasks.Tasks {
		if !task.Done {
			pendingTasks = append(pendingTasks, task)
		}
	}

	pendingTaskList := TaskList{Tasks: pendingTasks}

	err := taskTemplate.Execute(w, pendingTaskList)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

// ... (rest of the code remains the same)
