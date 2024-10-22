package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Customer represents a customer record in the CRM system.
type Customer struct {
	ID       int
	Name     string
	Email    string
	JoinDate time.Time
	// Add other customer attributes as needed
}

// CustomerStore represents a mock data store for customers using slice for data retrieval optimization
type CustomerStore struct {
	customers []Customer
	indexByID map[int]int // Index to find customer by ID efficiently
}

// NewCustomerStore creates a new CustomerStore
func NewCustomerStore() *CustomerStore {
	return &CustomerStore{
		customers: []Customer{},
		indexByID: make(map[int]int),
	}
}

// AddCustomer adds a customer to the store
func (cs *CustomerStore) AddCustomer(customer Customer) {
	cs.customers = append(cs.customers, customer)
	cs.indexByID[customer.ID] = len(cs.customers) - 1
}

// GetCustomerByID retrieves a customer from the store by its ID using the index
func (cs *CustomerStore) GetCustomerByID(customerID int) (Customer, bool) {
	idx, found := cs.indexByID[customerID]
	if !found {
		return Customer{}, false
	}
	return cs.customers[idx], true
}

// GetCustomersByName retrieves customers by a given name (partial match supported)
func (cs *CustomerStore) GetCustomersByName(name string) []Customer {
	var results []Customer

	for _, customer := range cs.customers {
		if len(name) == 0 || customer.Name == name || strings.Contains(customer.Name, name) {
			results = append(results, customer)
		}
	}

	return results
}

// GetCustomersByJoinDateRange retrieves customers within a specified date range
func (cs *CustomerStore) GetCustomersByJoinDateRange(startDate, endDate time.Time) []Customer {
	var results []Customer

	for _, customer := range cs.customers {
		if customer.JoinDate.After(startDate) && customer.JoinDate.Before(endDate) {
			results = append(results, customer)
		}
	}

	return results
}

func main() {
	// Initialize the customer store
	cs := NewCustomerStore()

	// Generate mock customers and add them to the store
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 10000; i++ {
		customer := Customer{
			ID:       i,
			Name:     fmt.Sprintf("Customer %d", i),
			Email:    fmt.Sprintf("customer%d@example.com", i),
			JoinDate: time.Now().AddDate(0, 0, -rand.Intn(365)),
		}
		cs.AddCustomer(customer)
	}

	// Optimized data retrieval strategies in action:

	// 1. Retrieve customer by ID
	customerIDToFind := 5000
	start := time.Now()
	customer, found := cs.GetCustomerByID(customerIDToFind)
	if found {
		fmt.Printf("Customer Found: %v\n", customer)
	} else {
		fmt.Printf("Customer with ID %d not found.\n", customerIDToFind)
	}
	fmt.Println("Time taken to retrieve customer by ID:", time.Since(start))

	// 2. Retrieve customers by name (partial match)
	nameToFind := "Customer 3"
	start = time.Now()
	customers := cs.GetCustomersByName(nameToFind)
	fmt.Printf("Customers matching '%s': %d\n", nameToFind, len(customers))
	fmt.Println("Time taken to retrieve customers by name:", time.Since(start))

	// 3. Retrieve customers by join date range
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()
	start = time.Now()
	customers = cs.GetCustomersByJoinDateRange(startDate, endDate)
	fmt.Printf("Customers joined between %s and %s: %d\n", startDate, endDate, len(customers))
	fmt.Println("Time taken to retrieve customers by join date range:", time.Since(start))
}
