package main  
import (  
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// ToDoItem represents a single to-do item
type ToDoItem struct {
    ID   int    `json:"id"`
    Task string `json:"task"`
}

var toDoItems []ToDoItem

func getAllToDoItems(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(toDoItems)
}

func getToDoItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
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

func createToDoItem(w http.ResponseWriter, r *http.Request) {
    var newItem ToDoItem
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newItem.ID = len(toDoItems) + 1
    toDoItems = append(toDoItems, newItem)

    json.NewEncoder(w).Encode(newItem)
}

func updateToDoItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var updatedItem ToDoItem
    if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, item := range toDoItems {
        if item.ID == id {
            updatedItem.ID = item.ID
            toDoItems[i] = updatedItem
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }

    http.Error(w, "To-do item not found", http.StatusNotFound)
}

func deleteToDoItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, item := range toDoItems {
        if item.ID == id {
            toDoItems = append(toDoItems[:i], toDoItems[i+1:]...)
            return
        }
    }

    http.Error(w, "To-do item not found", http.StatusNotFound)
}

func main() {
    r := mux.NewRouter()