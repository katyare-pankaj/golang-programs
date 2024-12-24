package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
	ID       int
	Name     string
	Quantity int
}

// Inventory is a slice to manage inventory items
type Inventory []InventoryItem

// AddItem adds a new item to the inventory
func (i *Inventory) AddItem(name string, quantity int) {
	newItem := InventoryItem{
		ID:       len(*i) + 1,
		Name:     name,
		Quantity: quantity,
	}
	*i = append(*i, newItem)
}

// ListItems prints the current inventory
func (i Inventory) ListItems() {
	if len(i) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("Inventory:")
	for _, item := range i {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", item.ID, item.Name, item.Quantity)
	}
}

// RemoveItemByID removes an item from the inventory by ID
func (i *Inventory) RemoveItemByID(id int) {
	for i, item := range *i {
		if item.ID == id {
			*i = append((*i)[:i], (*i)[i+1:]...)
			fmt.Printf("Item ID %d removed.\n", id)
			return
		}
	}
	fmt.Printf("Item ID %d not found.\n", id)
}

func main() {
	inventory := Inventory{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Inventory Tracking System!")
	for {
		fmt.Print("\nEnter command (add, list, remove, exit): ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Print("Enter item name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter item quantity: ")
			scanner.Scan()
			quantity, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid quantity. Please enter a number.")
				continue
			}

			inventory.AddItem(name, quantity)
			fmt.Printf("Item '%s' added to inventory.\n", name)

		case "list":
			inventory.ListItems()

		case "remove":
			fmt.Print("Enter item ID to remove: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			inventory.RemoveItemByID(id)

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}
