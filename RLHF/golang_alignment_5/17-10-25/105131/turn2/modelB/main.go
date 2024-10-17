// tasks.go
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Task represents a single task in the application
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Reusable function to handle errors
func handleError(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
	log.Println(err)
}

// nextAvailableID returns the next available ID for a new task
var nextAvailableID = 1

func getNextAvailableID() int {
	id := nextAvailableID
	nextAvailableID++
	return id
}

// TaskHandler serves the task-related API endpoints
type TaskHandler struct {
	tasks []Task
}

// Reusable function to find the task index by ID
func (th *TaskHandler) findTaskIndexByID(id int) int {
	for i, task := range th.tasks {
		if task.ID == id {
			return i
		}
	}
	return -1
}

// ValidateTask checks if a task is valid and returns an error if any field is missing or empty
func ValidateTask(task Task) error {
	if task.Title == "" {
		return errors.New("Title is required")
	}
	if task.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}

func (th *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	task.ID = getNextAvailableID()
	if err := ValidateTask(task); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	th.tasks = append(th.tasks, task)
	w.WriteHeader(http.StatusCreated)
}

func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Validate and update task logic using reusable functions
	// (Implementation details omitted)
}
