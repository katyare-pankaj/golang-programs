package main

import (
	"fmt"
)

// Product represents a product with a name and price.
type Product struct {
	Name  string
	Price float64
}

func main() {
	// Sample list of products
	products := []Product{
		{"Laptop", 1299.99},
		{"Mouse", 24.99},
		{"Keyboard", 49.99},
		{"Monitor", 199.99},
		{"Headphones", 89.99},
	}

	// Filter function to check if the price is greater than 50
	filterGreaterThanFifty := func(p Product) bool {
		return p.Price > 50
	}

	// Filter the products manually
	filteredProductsManually := filterProductsManually(products, filterGreaterThanFifty)
	fmt.Println("Filtered products manually:")
	printProducts(filteredProductsManually)

	// Filter the products using Go's built-in filter function
	filteredProductsBuiltin := filterProductsBuiltin(products, filterGreaterThanFifty)
	fmt.Println("\nFiltered products using Go's built-in filter:")
	printProducts(filteredProductsBuiltin)
}

// Filter the slice manually and return a new slice
func filterProductsManually(products []Product, condition func(Product) bool) []Product {
	filtered := []Product{}
	for _, product := range products {
		if condition(product) {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// Filter the slice using Go's built-in filter function
func filterProductsBuiltin(products []Product, condition func(Product) bool) []Product {
	return filter(products, condition)
}

// Filter function that uses a for loop to filter the slice
func filter(products []Product, condition func(Product) bool) []Product {
	return []Product{
		product for product := range products if condition(product)
	}
}

// Print the list of products
func printProducts(products []Product) {
	for _, product := range products {
		fmt.Printf("%v ($%.2f)\n", product.Name, product.Price)
	}
}