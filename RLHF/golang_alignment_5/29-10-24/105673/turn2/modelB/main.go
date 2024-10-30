package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// Product struct represents a product in the store
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Categories  []string
}

// InMemoryProductStore uses in-memory caching to store and retrieve products
type InMemoryProductStore struct {
	cache *cache.Cache
	mu    sync.Mutex
}

// NewInMemoryProductStore creates a new InMemoryProductStore
func NewInMemoryProductStore() *InMemoryProductStore {
	return &InMemoryProductStore{
		cache: cache.New(5*time.Minute, 10*time.Minute), // Cache items for 5 minutes, clean up every 10 minutes
	}
}

// StoreProduct stores a product in the in-memory store
func (ims *InMemoryProductStore) StoreProduct(product *Product) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	ims.cache.Set(fmt.Sprint(product.ID), product, cache.DefaultExpiration)
}

// GetProduct retrieves a product from the in-memory store based on its ID
func (ims *InMemoryProductStore) GetProduct(id int) (*Product, bool) {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	val, found := ims.cache.Get(fmt.Sprint(id))
	if found {
		return val.(*Product), true
	}
	return nil, false
}

// getRandomProduct fetches a random product from the database (in a real scenario)
func getRandomProduct() *Product {
	// In a real application, this would query the database to fetch the product.
	time.Sleep(time.Millisecond * 50) // Simulate database query time
	productID := rand.Intn(1000)
	return &Product{
		ID:          productID,
		Name:        fmt.Sprintf("Product %d", productID),
		Description: fmt.Sprintf("This is product %d.", productID),
		Price:       float64(rand.Intn(1000)) + 0.99,
		Categories:  []string{"Electronics", "Smart Devices", "Gaming"},
	}
}

// searchProducts simulates a real-time search for products based on a given query
func searchProducts(query string, store *InMemoryProductStore) []*Product {
	// In a real application, this would perform a search in the database or elasticsearch
	// Based on the query, we'll fetch some random products here.
	results := []*Product{}
	for i := 0; i < 5; i++ {
		product := getRandomProduct()
		// Check if the product is already in the cache to avoid redundant queries
		if _, cached := store.GetProduct(product.ID); !cached {
			store.StoreProduct(product)
		}
		results = append(results, product)
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	store := NewInMemoryProductStore()

	// Simulate a real-time search query
	query := "smartphone"
	startTime := time.Now()
	searchResults := searchProducts(query, store)
	endTime := time.Now()

	// Output the search results and cache hit statistics
	fmt.Println("Search Results:")
	for _, product := range searchResults {
		fmt.Printf("Product %d: %s ($%.2f)\n", product.ID, product.Name, product.Price)
	}
	fmt.Println("------------------------")
	fmt.Printf("Cache Hits: %d\n", store.cache.ItemCount())