package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handle getting an item by ID
func getItem(w http.ResponseWriter, r *http.Request) {
	// Extract the "id" parameter from the URL
	itemIDStr := r.URL.Query().Get("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %v", err), http.StatusBadRequest)
		return
	}

	item := Item{ID: itemID, Name: "Sample Item"}
	response, err := json.Marshal(item)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal response: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Handle creating a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// In a real application, save the item to a database or data store

	response, err := json.Marshal(item)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal response: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func main() {
	// Register handlers for different routes
	http.HandleFunc("/items", createItem) // POST /items
	http.HandleFunc("/items/", getItem)   // GET /items/{id}

	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
