package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/btree"
)

// Customer represents a customer record in the CRM system.
type Customer struct {
	ID       int
	Name     string
	Email    string
	JoinDate time.Time
	// Add other customer attributes as needed
}

// CustomerStore represents a mock data store for customers using a slice for data retrieval optimization.
type CustomerStore struct {
	customers []Customer
	indexByID map[int]int // Index to find customer by ID efficiently
	index     *CustomerIndex
	mu        sync.RWMutex // Mutex to protect concurrent access
}

// CustomerIndex is an in-memory B-tree index for customers based on their ID.
type CustomerIndex struct {
	Tree *btree.BTree
	m    sync.RWMutex
}

// NewCustomerIndex creates a new CustomerIndex.
func NewCustomerIndex() *CustomerIndex {
	return &CustomerIndex{
		Tree: btree.New(32), // Adjust the degree based on your use case.
	}
}

// AddCustomerIndex adds an entry to the customer index.
func (ci *CustomerIndex) AddCustomerIndex(customer Customer) {
	ci.m.Lock()
	defer ci.m.Unlock()
	ci.Tree.ReplaceOrInsert(btree.Int(customer.ID))
}

// GetCustomerByID retrieves a customer from the index by ID.
func (ci *CustomerIndex) GetCustomerByID(customerID int) (*Customer, bool) {
	ci.m.RLock()
	defer ci.m.RUnlock()
	it := ci.Tree.Get(btree.Int(customerID))
	if it != nil {
		return nil, true
	}
	return nil, false
}

// NewCustomerStore creates a new CustomerStore.
func NewCustomerStore() *CustomerStore {
	return &CustomerStore{
		customers: []Customer{},
		indexByID: make(map[int]int),
		index:     NewCustomerIndex(),
	}
}

// AddCustomer adds a new customer to the store.
func (cs *CustomerStore) AddCustomer(customer Customer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.customers = append(cs.customers, customer)
	cs.indexByID[customer.ID] = len(cs.customers) - 1
	cs.index.AddCustomerIndex(customer)
}

// GetCustomerByID retrieves a customer from the store by ID.
func (cs *CustomerStore) GetCustomerByID(customerID int) (*Customer, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	if index, found := cs.indexByID[customerID]; found {
		return &cs.customers[index], true
	}
	return nil, false
}

// AggregateCustomersByDate aggregates the number of customers joined on each day.
func (cs *CustomerStore) AggregateCustomersByDate() map[time.Time]int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	aggregatedData := make(map[time.Time]int)
	for _, customer := range cs.customers {
		joinDate := customer.JoinDate.Truncate(24 * time.Hour)
		aggregatedData[joinDate]++
	}

	return aggregatedData
}

func main() {
	// Initialize the customer store
	cs := NewCustomerStore()

	// Generate mock customers and add them to the store
	for i := 1; i <= 1000; i++ {
		customer := Customer{
			ID:       i,
			Name:     fmt.Sprintf("Customer %d", i),
			Email:    fmt.Sprintf("customer%d@example.com", i),
			JoinDate: time.Now().AddDate(0, 0, -i%365), // Random join date within the past year
		}
		cs.AddCustomer(customer)
	}

	// Retrieve a customer by ID
	if customer, found := cs.GetCustomerByID(500); found {
		fmt.Printf("Retrieved Customer: %+v\n", customer)
	} else {
		fmt.Println("Customer not found.")
	}

	// Aggregate customers by join date
	aggregatedData := cs.AggregateCustomersByDate()
	fmt.Println("Aggregated Customers by Join Date:")
	for date, count := range aggregatedData {
		fmt.Printf("%s: %d\n", date.Format("2006-01-02"), count)
	}
}
