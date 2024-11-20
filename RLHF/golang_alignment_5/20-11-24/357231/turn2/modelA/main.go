package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ToDoItem struct represents a single to-do item
type ToDoItem struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var toDoItems []ToDoItem

// IndexHandler returns a list of all to-do items
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(toDoItems)
}

// CreateHandler adds a new to-do item
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var item ToDoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	item.ID = len(toDoItems) + 1
	toDoItems = append(toDoItems, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// ShowHandler returns a single to-do item by its ID
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, item := range toDoItems {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "To-do item not found", http.StatusNotFound)
}

// UpdateHandler updates an existing to-do item
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var item ToDoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for i, storedItem := range toDoItems {
		if storedItem.ID == id {
			toDoItems[i] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "To-do item not found", http.StatusNotFound)
}

// DeleteHandler deletes a to-do item by its ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error converting ID to integer: %v\n", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, item := range toDoItems {
		if item.ID == id {
			toDoItems = append(toDoItems[:i], toDoItems[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "To-do item not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", IndexHandler).Methods("GET")
	r.HandleFunc("/todos", CreateHandler).Methods("POST")
	r.HandleFunc("/todos", ShowHandler).Queries("id", "{id:[0-9]+}").Methods("GET")
	r.HandleFunc("/todos", UpdateHandler).Queries("id", "{id:[0-9]+}").Methods("PUT")
	r.HandleFunc("/todos", DeleteHandler).Queries("id", "{id:[0-9]+}").Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
