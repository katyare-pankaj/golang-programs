package main

import (
	"fmt"
)

// Item struct
type Item struct {
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

func main() {
	showMenu()
}

// showMenu presents the main menu options to the user.
func showMenu() {
	for {
		fmt.Println(`
        1. Add Item
        2. Remove Item
        3. Display Inventory
        4. Update Item Quantity
        5. Exit
        Enter your choice: `)
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			addItem()
		case 2:
			removeItem()
		case 3:
			displayInventory()
		case 4:
			updateItemQuantity()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select between 1-5.")
		}
	}
}

// addItem adds a new item to the inventory slice.
func addItem() {
	var item Item
	fmt.Print("Item Name: ")
	fmt.Scanln(&item.Name)
	fmt.Print("Item Quantity: ")
	fmt.Scanln(&item.Quantity)
	fmt.Print("Item Price: ")
	fmt.Scanln(&item.Price)
	inventory = append(inventory, item)
	fmt.Println("Item added successfully.")
}

// removeItem removes an item from the inventory by name.
func removeItem() {
	var itemName string
	fmt.Print("Item Name to Remove: ")
	fmt.Scanln(&itemName)
	for index, item := range inventory {
		if item.Name == itemName {
			inventory = append(inventory[:index], inventory[index+1:]...)
			fmt.Println("Item removed successfully.")
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}

// displayInventory lists all the items in the inventory with their details.
func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}
	fmt.Println("\nInventory:")
	fmt.Println("--------------------")
	for _, item := range inventory {
		fmt.Printf("%s: %d units @ $%.2f each\n", item.Name, item.Quantity, item.Price)
	}
	fmt.Println("--------------------")
}

// updateItemQuantity updates the quantity of an item in the inventory by name.
func updateItemQuantity() {
	var itemName string
	fmt.Print("Item Name to Update: ")
	fmt.Scanln(&itemName)
	for index, item := range inventory {
		if item.Name == itemName {
			var newQuantity int
			fmt.Print("Enter new Quantity: ")
			fmt.Scanln(&newQuantity)
			inventory[index].Quantity = newQuantity
			fmt.Println("Item quantity updated successfully.")
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}
