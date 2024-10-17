package main

import (
	"fmt"
)

type Order struct {
	CustomerName string
	OrderID      string
	Items        []Item
}

type Item struct {
	ProductName string
	Quantity    int
	Price       float32
}

func transformOrders(orders []Order) []Order {
	transformedOrders := make([]Order, 0)
	for _, order := range orders {
		transformedOrder := order
		for i, item := range order.Items {
			transformedOrder.Items[i].Price = item.Price * float32(item.Quantity)
		}
		transformedOrders = append(transformedOrders, transformedOrder)
	}
	return transformedOrders
}

func main() {
	orders := []Order{
		{CustomerName: "John Doe", OrderID: "1", Items: []Item{{ProductName: "Apple", Quantity: 2, Price: 2.5}}},
		{CustomerName: "Jane Doe", OrderID: "2", Items: []Item{{ProductName: "Banana", Quantity: 3, Price: 1.8}}},
	}
	transformedOrders := transformOrders(orders)
	for _, order := range transformedOrders {
		fmt.Printf("Customer Name: %s, Order ID: %s\n", order.CustomerName, order.OrderID)
		for _, item := range order.Items {
			fmt.Printf("Product Name: %s, Quantity: %d, Total Price: %.2f\n", item.ProductName, item.Quantity, item.Price)
		}
		fmt.Println()
	}
}
