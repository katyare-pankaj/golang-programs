package main

import (
	"fmt"
	"sort"
)

// Customer represents a customer record
type Customer struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Original slice of customer records (immutable)
	customers := []Customer{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Charlie", Age: 22},
		{ID: 4, Name: "David", Age: 35},
		{ID: 5, Name: "Emma", Age: 28},
	}

	// 1. Filter customers by age
	filteredCustomers := filterCustomersByAge(customers, 28, 35)
	fmt.Println("\nFiltered customers (age 28 to 35):")
	displayCustomers(filteredCustomers)

	// 2. Aggregate customers' ages
	totalAge, averageAge := aggregateAges(customers)
	fmt.Println("\nTotal age of all customers:", totalAge)
	fmt.Println("Average age of all customers:", averageAge)

	// 3. Sort customers by age
	sortedCustomers := sortCustomersByAge(customers)
	fmt.Println("\nSorted customers by age:")
	displayCustomers(sortedCustomers)
}

// filterCustomersByAge returns a new slice containing customers with ages within the specified range.
func filterCustomersByAge(customers []Customer, minAge, maxAge int) []Customer {
	var filteredCustomers []Customer
	for _, customer := range customers {
		if customer.Age >= minAge && customer.Age <= maxAge {
			filteredCustomers = append(filteredCustomers, customer)
		}
	}
	return filteredCustomers
}

// aggregateAges returns the total and average ages of all customers.
func aggregateAges(customers []Customer) (totalAge int, averageAge float64) {
	var totalAge int
	for _, customer := range customers {
		totalAge += customer.Age
	}
	averageAge = float64(totalAge) / float64(len(customers))
	return totalAge, averageAge
}

// sortCustomersByAge returns a new slice containing customers sorted by age.
func sortCustomersByAge(customers []Customer) []Customer {
	// We'll create a copy of the original slice to avoid modifying it.
	sortedCustomers := make([]Customer, len(customers))
	copy(sortedCustomers, customers)

	// Sort the copied slice using a custom sort function.
	sort.Slice(sortedCustomers, func(i, j int) bool {
		return sortedCustomers[i].Age < sortedCustomers[j].Age
	})

	return sortedCustomers
}

// displayCustomers displays the list of customers.
func displayCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("Customer ID: %d, Name: %s, Age: %d\n", customer.ID, customer.Name, customer.Age)
	}
}
