package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Inventory represents the inventory of products
type Inventory struct {
	Items map[string]int `json:"items"`
	lock  sync.RWMutex   `json:"-"`
}

// LoadInventory loads inventory data from a file
func LoadInventory(filename string) *Inventory {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var inv Inventory
	if err := json.Unmarshal(data, &inv); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return &inv
}

// AddItem adds an item to the inventory
func (inv *Inventory) AddItem(item string, quantity int) {
	inv.lock.Lock()
	defer inv.lock.Unlock()

	inv.Items[item] += quantity
}

// RemoveItem removes an item from the inventory
func (inv *Inventory) RemoveItem(item string, quantity int) bool {
	inv.lock.Lock()
	defer inv.lock.Unlock()

	currentQuantity, ok := inv.Items[item]
	if !ok {
		return false
	}

	if quantity > currentQuantity {
		return false
	}

	inv.Items[item] -= quantity
	if inv.Items[item] == 0 {
		delete(inv.Items, item)
	}
	return true
}

// UpdateItem updates the quantity of an item in the inventory
func (inv *Inventory) UpdateItem(item string, quantity int) {
	inv.lock.Lock()
	defer inv.lock.Unlock()

	inv.Items[item] = quantity
}

// GetItemCount returns the count of a specific item in the inventory
func (inv *Inventory) GetItemCount(item string) int {
	inv.lock.RLock()
	defer inv.lock.RUnlock()

	return inv.Items[item]
}

func handleAddItem(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	quantityStr := r.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	inv.AddItem(item, quantity)
	fmt.Fprintf(w, "Item added successfully!")
}

func handleRemoveItem(w http.ResponseWriter, r *http.Request) {
	// Implement similar logic for remove item
}

func handleUpdateItem(w http.ResponseWriter, r *http.Request) {
	// Implement similar logic for update item
}

func handleGetItemCount(w http.ResponseWriter, r *http.Request) {
	// Implement similar logic for get item count
}

var inv *Inventory

func main() {
	inv = LoadInventory("inventory.json")
	http.HandleFunc("/add", handleAddItem)
	http.HandleFunc("/remove", handleRemoveItem)
	http.HandleFunc("/update", handleUpdateItem)
	http.HandleFunc("/count", handleGetItemCount)

	fmt.Println("Inventory management system running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
