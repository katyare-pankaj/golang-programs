package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// Property struct represents a property in the system
type Property struct {
	ID        int
	Address   string
	Price     int
	Amenities []string
}

// InMemoryPropertyStore uses in-memory caching to store and retrieve properties
type InMemoryPropertyStore struct {
	cache *cache.Cache
	mu    sync.Mutex
}

// NewInMemoryPropertyStore creates a new InMemoryPropertyStore
func NewInMemoryPropertyStore() *InMemoryPropertyStore {
	return &InMemoryPropertyStore{
		cache: cache.New(5*time.Minute, 10*time.Minute), // Cache items for 5 minutes, clean up every 10 minutes
	}
}

// StoreProperty stores a property in the in-memory store
func (ims *InMemoryPropertyStore) StoreProperty(property *Property) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	ims.cache.Set(fmt.Sprint(property.ID), property, cache.DefaultExpiration)
}

// GetProperty retrieves a property from the in-memory store based on its ID
func (ims *InMemoryPropertyStore) GetProperty(id int) (*Property, bool) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	val, found := ims.cache.Get(fmt.Sprint(id))
	if found {
		return val.(*Property), true
	}
	return nil, false
}

// Example usage:
func main() {
	store := NewInMemoryPropertyStore()

	// Store some properties in the in-memory store
	store.StoreProperty(&Property{
		ID:        1,
		Address:   "123 Main Street, Anytown, USA",
		Price:     200000,
		Amenities: []string{"Pool", "Garage"},
	})

	store.StoreProperty(&Property{
		ID:        2,
		Address:   "456 Park Avenue, Cityville, USA",
		Price:     350000,
		Amenities: []string{"Gym", "Balcony"},
	})

	// Retrieve a property by ID
	idToFind := 1
	prop, found := store.GetProperty(idToFind)
	if found {
		fmt.Printf("Property Found: %+v\n", prop)
	} else {
		fmt.Printf("Property with ID %d not found.\n", idToFind)
	}

	// Output: Property Found: &{ID:1 Address:123 Main Street, Anytown, USA Price:200000 Amenities:[Pool Garage]}
}
