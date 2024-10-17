package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
}

var tasks []Task

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasksDueToday).Methods("GET")

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func getTasksDueToday(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	today := time.Now().Truncate(24 * time.Hour)
	var dueTodayTasks []Task

	for _, task := range tasks {
		if task.DueDate.Truncate(24*time.Hour) == today {
			dueTodayTasks = append(dueTodayTasks, task)
		}
	}

	err := json.NewEncoder(w).Encode(dueTodayTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
