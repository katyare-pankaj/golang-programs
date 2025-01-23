package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Task struct {
	Title       string
	Description string
	Status      string // e.g., "Pending", "In Progress", "Completed"
}

type TaskList struct {
	Tasks []Task
}

func main() {
	// Initialize an empty task list
	taskList := &TaskList{}

	// Add tasks to the list
	taskList.AddTask("Task 1", "Description of Task 1", "Pending")
	taskList.AddTask("Task 2", "Description of Task 2", "In Progress")

	// Update a task status
	taskList.UpdateTaskStatus(0, "Completed") // Update Task 1 to "Completed"

	// Render the task list to an HTML file
	if err := taskList.RenderHTML("task_list.html"); err != nil {
		log.Fatalf("Error rendering HTML: %v", err)
	}

	fmt.Println("Task list rendered successfully!")
}

// AddTask adds a new task to the task list
func (tl *TaskList) AddTask(title, description, status string) {
	tl.Tasks = append(tl.Tasks, Task{Title: title, Description: description, Status: status})
}

// UpdateTaskStatus updates the status of a task given its index
func (tl *TaskList) UpdateTaskStatus(index int, status string) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Status = status
	} else {
		fmt.Println("Invalid task index")
	}
}

// RenderHTML renders the task list to an HTML file
func (tl *TaskList) RenderHTML(filename string) error {
	const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Task List</title>
</head>
<body>
    <h1>Task List</h1>
    <ul>
    {{ range .Tasks }}
        <li>
            <strong>{{ .Title }}</strong>: {{ .Description }} (Status: {{ .Status }})
        </li>
    {{ end }}
    </ul>
</body>
</html>`

	t, err := template.New("tasklist").Parse(tpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	if err := t.Execute(file, tl); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}
