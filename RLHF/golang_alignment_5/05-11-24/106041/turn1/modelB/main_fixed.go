package main

import (
	"encoding/json"
	"fmt"
)

// Struct definitions
type Order struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
}

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order *Order `json:"-"` // Ignore this field to break the cycle
}

// Revised structs to break the cycle
type OrderV2 struct {
	ID       int        `json:"id"`
	Customer CustomerV2 `json:"customer"`
}

type CustomerV2 struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	OrderID int    `json:"order_id,omitempty"` // Use OrderID instead of a full Order object
}

func main() {
	// Original structs with a reference cycle
	customer := Customer{
		ID:   1,
		Name: "John Doe",
	}
	customer.Order = &Order{
		ID:       100,
		Customer: customer,
	}

	// Attempting to marshal original customer (will omit the Order field)
	jsonData, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error marshalling customer with cycle:", err)
		return
	}
	fmt.Println("JSON data before breaking cycle:", string(jsonData))

	// Revised structs without a reference cycle
	order := OrderV2{
		ID: 100,
		Customer: CustomerV2{
			ID:      customer.ID,
			Name:    customer.Name,
			OrderID: 100,
		},
	}

	// Marshaling after breaking cycle
	jsonDataV2, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshalling order after breaking cycle:", err)
		return
	}
	fmt.Println("JSON data after breaking cycle:", string(jsonDataV2))
}
