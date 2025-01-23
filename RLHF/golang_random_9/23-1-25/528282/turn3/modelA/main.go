package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskList struct {
	Tasks []Task
}

var (
	tasks        TaskList
	taskTemplate *template.Template
)

func init() {
	// Initialize the task list
	tasks = TaskList{
		Tasks: []Task{
			{ID: 1, Title: "Buy milk", Done: false},
			{ID: 2, Title: "Learn Go", Done: true},
			{ID: 3, Title: "Write a blog post", Done: false},
		},
	}

	// Parse the task list template
	taskTemplate, _ = template.New("tasks").Parse(`
        <html>
            <head>
                <title>Task List</title>
            </head>
            <body>
                <h1>Pending Tasks</h1>
                <ul>
                {{ range .Tasks }}
                    <li>
                        <input type="checkbox" {{ if .Done }} checked {{ end }}>
                        {{ .Title }}
                        <a href="/update?id={{ .ID }}">Edit</a>
                        <a href="/delete?id={{ .ID }}">Delete</a>
                    </li>
                {{ end }}
                </ul>
                <form action="/add" method="post">
                    <input type="text" name="title" placeholder="New task...">
                    <input type="submit" value="Add">
                </form>
            </body>
        </html>
    `)
}

func main() {
	http.HandleFunc("/", showTasks)
	http.HandleFunc("/add", addTask)
	http.HandleFunc("/update", updateTask)
	http.HandleFunc("/delete", deleteTask)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	// Filter tasks to show only pending ones
	pendingTasks := TaskList{}
	for _, task := range tasks.Tasks {
		if !task.Done {
			pendingTasks.Tasks = append(pendingTasks.Tasks, task)
		}
	}

	err := taskTemplate.Execute(w, pendingTasks)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		if title != "" {
			newTask := Task{ID: len(tasks.Tasks) + 1, Title: title, Done: false}
			tasks.Tasks = append(tasks.Tasks, newTask)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for i, task := range tasks.Tasks {
			if task.ID == id {
				tasks.Tasks[i].Done = !task.Done
				break
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for i, task := range tasks.Tasks {
			if task.ID == id {
				tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
				break
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
