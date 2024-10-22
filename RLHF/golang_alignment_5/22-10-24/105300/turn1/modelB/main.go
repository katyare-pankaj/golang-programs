package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Customer represents a customer record in the CRM system.
type Customer struct {
	ID       int
	Name     string
	JoinDate time.Time
}

// CustomerStore represents a mock data store for customers.
type CustomerStore map[int]Customer

// ShardedCustomerStore represents a sharded customer store with multiple shards.
type ShardedCustomerStore struct {
	Shards    []CustomerStore
	ShardSize int
}

// NewShardedCustomerStore creates a new sharded customer store with the specified number of shards.
func NewShardedCustomerStore(numShards int) *ShardedCustomerStore {
	shards := make([]CustomerStore, numShards)
	// Initialize each shard
	for i := range shards {
		shards[i] = make(CustomerStore)
	}
	return &ShardedCustomerStore{
		Shards:    shards,
		ShardSize: 100, // You can adjust this based on your use case
	}
}

// getShardIndex returns the index of the shard for a given customer ID.
func (scs *ShardedCustomerStore) getShardIndex(customerID int) int {
	return customerID % len(scs.Shards) // Use modulo to get a valid index
}

// AddCustomer adds a customer to the sharded customer store.
func (scs *ShardedCustomerStore) AddCustomer(customer Customer) {
	shardIndex := scs.getShardIndex(customer.ID)
	scs.Shards[shardIndex][customer.ID] = customer
}

// GetCustomer retrieves a customer from the sharded customer store based on its ID.
func (scs *ShardedCustomerStore) GetCustomer(customerID int) (Customer, bool) {
	shardIndex := scs.getShardIndex(customerID)
	customer, exists := scs.Shards[shardIndex][customerID]
	return customer, exists
}

// AggregateCustomersByDate aggregates the number of customers joined on each day.
func (scs *ShardedCustomerStore) AggregateCustomersByDate() map[time.Time]int {
	aggregatedData := make(map[time.Time]int)

	// Iterate through each shard and aggregate the data
	for _, shard := range scs.Shards {
		for _, customer := range shard {
			joinDate := customer.JoinDate.Truncate(24 * time.Hour)
			aggregatedData[joinDate]++
		}
	}

	return aggregatedData
}

func main() {
	// Initialize the sharded customer store with 5 shards
	scs := NewShardedCustomerStore(5)

	// Generate mock customers and add them to the sharded store
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 1000; i++ {
		customer := Customer{
			ID:       i,
			Name:     fmt.Sprintf("Customer %d", i),
			JoinDate: time.Now().AddDate(0, 0, -rand.Intn(365)),
		}
		scs.AddCustomer(customer)
	}

	// Example of aggregating customers by join date
	aggregatedData := scs.AggregateCustomersByDate()
	fmt.Println("Aggregated Customers by Join Date:")
	for date, count := range aggregatedData {
		fmt.Printf("%s: %d\n", date.Format("2006-01-02"), count)
	}
}
