package main

import (
	"fmt"
)

// Product struct
type Product struct {
	Name  string
	Price float64
}

func main() {
	// Initialize a slice of Product structs
	products := []Product{
		{Name: "Apple", Price: 2.5},
		{Name: "Banana", Price: 1.8},
		{Name: "Orange", Price: 3.0},
		{Name: "Grapes", Price: 4.0},
		{Name: "Mango", Price: 2.0},
	}

	minPrice := 2.5 // Define the minimum price filter

	// Filter the products slice to include only products with price greater than minPrice
	filteredProducts := filterProducts(products, func(p Product) bool {
		return p.Price > minPrice
	})

	// Print the filtered products
	fmt.Println("Filtered Products:")
	for _, product := range filteredProducts {
		fmt.Printf("%s: %.2f\n", product.Name, product.Price)
	}
}

// filterProducts function takes a slice of Product structs and a predicate function as input
// and returns a new slice containing only those structs for which the predicate function returns true.
func filterProducts(products []Product, predicate func(Product) bool) []Product {
	var filtered []Product
	for _, product := range products {
		if predicate(product) {
			filtered = append(filtered, product)
		}
	}
	return filtered
}
